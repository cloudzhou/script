import re
import os
import json
import random
import socket
import threading
import traceback
from urlparse import urlparse, parse_qs
from datetime import datetime
from BaseHTTPServer import BaseHTTPRequestHandler, HTTPServer
from SocketServer import BaseRequestHandler, ThreadingTCPServer

DIRNAME = os.path.dirname(os.path.abspath(__file__))

GATE = {}

DELIVER = {}

BINS = {'user1.bin': os.path.join(DIRNAME, 'user1.bin'), 'user2.bin': os.path.join(DIRNAME, 'user2.bin')}
HTML = os.path.join(DIRNAME, 'local_postman_server.html')

DEVICES = [
    ['55e83b2d26f2266ffe4bc989f0a01fb55dbe52ce', '96c73f149ea9829065029e0df1918ed8761cd47e', {'productbatch_id': None, 'bssid': 'b18a0', 'ptype': 23701, 'activate_status': 1, 'serial': '3c1df254', 'id': 6, 'description': '839cc', 'last_active': '2015-05-19 20:13:03', 'rom_version': None, 'last_pull': '2014-06-13 17:06:16', 'last_push': '2014-06-13 17:06:16', 'location': '', 'metadata': 'cf00e', 'status': 1, 'updated': '2015-05-19 20:13:03', 'latest_rom_version': 'v1.1', 'activated_at': '2015-05-19 15:24:49', 'visibly': 1, 'is_private': 1, 'product_id': 3, 'name': '80c40', 'created': '2014-05-19 15:04:22', 'is_frozen': 0, 'key_id': 1018}],
    ['dffdb5f00d31d2856dce008ad63c7403ef3780f2', 'bad1cddfddd89d02de58bc4017c6f7f10771decb', {'productbatch_id': None, 'bssid': '', 'ptype': 23701, 'activate_status': 1, 'serial': 'a2d4b221', 'id': 7, 'description': 'description Of device(serial: d4bd3e38)', 'last_active': '2015-05-19 20:13:03', 'rom_version': None, 'last_pull': '2014-05-19 15:09:22', 'last_push': '2014-05-19 15:09:22', 'location': '', 'metadata': '', 'status': 1, 'updated': '2015-05-19 20:13:03', 'latest_rom_version': 'v1.1', 'activated_at': '2014-05-19 15:09:22', 'visibly': 1, 'is_private': 1, 'product_id': 3, 'name': 'switch002', 'created': '2014-05-19 15:09:22', 'is_frozen': 0, 'key_id': 1020}],
    ['1704244f8d9f162cddbbaec37b4edd3444bfebf8', 'f2160e452042487f2531187c8aafb668d06c6d4e', {'productbatch_id': None, 'bssid': '', 'ptype': 23701, 'activate_status': 1, 'serial': '99f33cc6', 'id': 8, 'description': 'description Of device(serial: 4c2a3aad)', 'last_active': '2015-05-19 20:13:03', 'rom_version': None, 'last_pull': '2014-05-19 15:14:56', 'last_push': '2014-05-19 15:14:56', 'location': '', 'metadata': '', 'status': 1, 'updated': '2015-05-19 20:13:03', 'latest_rom_version': 'v1.1', 'activated_at': '2014-05-19 15:14:56', 'visibly': 1, 'is_private': 1, 'product_id': 3, 'name': 'switch003', 'created': '2014-05-19 15:14:56', 'is_frozen': 0, 'key_id': 1022}],
]

def deliver_to_device(device_id, jsonobj):
    if device_id not in GATE:
        return False
    try:
        GATE[device_id].request.sendall(json.dumps(jsonobj) + '\n')
        return True
    except Exception, e:
        print e
        print traceback.format_exc()
    return False

class JsonHandler():

    def handle(self, postmanHandler, jsonobj):
        print 'handle postman: %s' % json.dumps(jsonobj)
        if 'path' not in jsonobj and 'nonce' in jsonobj:
            nonce = jsonobj['nonce']
            if nonce in DELIVER:
                event = DELIVER[nonce]
                event.data = {'status': 200, 'nonce': nonce, 'message': 'deliver to device success'}
                event.set()
            return
            
        path = jsonobj['path']
        token = ''
        if 'meta' in jsonobj and 'Authorization' in jsonobj['meta']:
            token = jsonobj['meta']['Authorization'].split(' ')[1]

        device = None
        for x in DEVICES:
            if token == x[0]:
                device = x[2]

        # ping
        if re.match('^/v1/ping/?$', path):
            now = datetime.now()
            nowstr = now.strftime('%Y-%m-%d %H:%M:%S')
            epoch = int(now.strftime('%s'))
            return {'datetime': nowstr, 'epoch': epoch, 'message': 'ping success'}
        # get device
        elif re.match('^/v1/device/?$', path):
            return {'device': device}
        # device activate
        elif re.match('^/v1/device/activate/?$', path):
            return {'device': device, 'key': {'token': token}, 'token': None}
        # device identify
        elif re.match('^/v1/device/identify/?$', path):
            if device:
                device_id = device['id']
                if device_id in GATE:
                    if postmanHandler != GATE[device_id]:
                        self.close_quiet(GATE[device_id])
                GATE[device_id] = postmanHandler
                return {'device': device}
        # device quit
        elif re.match('^/v1/device/quit/?$', path):
            return None

        # datapoint
        m = re.match('^/v1/datastreams/([a-zA-Z0-9_\+\-\.\:]+)/datapoint/?', path)
        if m:
            return jsonobj

        # rpc
        m = re.match('^/v1/device/rpc/?', path)
        if m:
            return jsonobj

        # else
        return {'status': 404, 'message': 'what is that?'}

    def close_quiet(self, postmanHandler):
        try:
            postmanHandler.request.sendall('{"status": 200, "message": "byebye, close by other peer login with same master key"}\n')
            postmanHandler.request.shutdown(socket.SHUT_RDWR)
            postmanHandler.request.close()
            postmanHandler.finish()
        except Exception, e:
            print e
            print traceback.format_exc()

class PostmanHandler(BaseRequestHandler):

    def handle(self):
        self.jsonHandler = JsonHandler()
        self._buffer = ''
        try:
            while True:
                recv_buffer = self.request.recv(4096)
                if not recv_buffer:
                    break
                r = recv_buffer.rfind('\n')
                if r < 0:
                    self._buffer = self._buffer + recv_buffer
                    if len(self._buffer) > 1024 * 1024:
                        break
                    if recv_buffer[-1] == '\x06':
                        break
                    continue
                line = self._buffer + recv_buffer[0:r]
                self._buffer = recv_buffer[r+1:]
                jsonstrs = line.split('\n')
                for jsonstr in jsonstrs:
                    jsonobj = json.loads(jsonstr)
                    r = self.jsonHandler.handle(self, jsonobj)
                    if not r:
                        break
                    if 'nonce' in jsonobj:
                        r['nonce'] = jsonobj['nonce']
                    if 'status' not in r:
                        r['status'] = 200
                    self.request.sendall(json.dumps(r) + '\n')
        except Exception, e:
            print e
            print traceback.format_exc()

class IotHttpHandler(BaseHTTPRequestHandler):

    def request_meta(self):
        up = urlparse(self.path)
        path = up.path
        meta = {}
        authorization = self.headers.getheader('Authorization')
        if authorization:
            meta['Authorization'] = authorization
        get = {}
        for x, y in parse_qs(up.query).items():
            get[x] = y[0]
        body = {}
        content_type = self.headers.getheader('Content-type')
        if content_type and content_type.startswith('application/x-www-form-urlencoded'):
            length = int(self.headers.getheader('content-length'))
            body = json.loads(self.rfile.read(length))
        token = ''
        if authorization:
            token = authorization.split(' ')[1]
        device = None
        for x in DEVICES:
            if token == x[1]:
                device = x[2]
        nonce = random.randint(0, 1<<20)
        return path, nonce, token, device, {'nonce': nonce, 'path': path, 'meta': meta, 'get': get, 'body': body}

    def do(self):
        path, nonce, token, device, jsonobj = self.request_meta()
        if re.match('^/?$', path):
            self.send_status_header('text/html')
            self.write_data(open(HTML, 'r').read())
        elif re.match('^/v1/device/rom/?$', path):
            filename = jsonobj['get']['filename']
            if filename not in BINS:
                self.send_status_header('application/json')
                self.write_data('{"status": 404, "nonce": %s, "message": "bin not found"}' % nonce)
                return
            filepath = BINS[filename]
            self.send_status_header('application/octet-stream')
            self.write_data(open(filepath, 'r').read())
            return
        self.send_status_header('application/json')
        if device and nonce:
            d = deliver_to_device(device['id'], jsonobj)
            if d:
                try:
                    event = threading.Event()
                    DELIVER[nonce] = event
                    w = event.wait(5)
                    if not w:
                        self.write_data('{"status": 500, "nonce": %s, "message": "remote device is busy or some errors occured"}' % nonce)
                        return
                    self.write_data(json.dumps(event.data))
                    return
                except Exception, e:
                    print e
                    print traceback.format_exc()
                finally:
                    del DELIVER[nonce]

        self.write_data('{"status": 404, "nonce": %s, "message": "deivce not found"}' % nonce)

    def do_GET(self):
        self.do()

    def do_POST(self):
        self.do()

    def send_status_header(self, content_type):
        self.send_response(200)
        self.send_header('Content-Type', content_type)

    def write_data(self, data):
        if not data:
            self.end_headers()
            return
        self.send_header('Content-Length', len(data))
        self.send_header('Connection', 'close')
        self.end_headers()
        self.wfile.write(data)

if __name__ == '__main__':
    httpServer = HTTPServer(('0.0.0.0', 8080), IotHttpHandler)
    httpServer.allow_reuse_address = True
    t = threading.Thread(target=httpServer.serve_forever)
    t.daemon = True
    t.start()
    postmanServer = ThreadingTCPServer(('0.0.0.0', 8000), PostmanHandler)
    postmanServer.allow_reuse_address = True
    postmanServer.serve_forever()

