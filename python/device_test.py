import os
import re
import sys
import json
import random
import socket

HOST = '114.215.177.97' # 'localhost'
PORT = 8000
TOKEN = '6af62e95120640d78a3cea447ba313832ca96a19' # '4340f23e0ceee9f7c550d76bd322d14296b24537'

# {"nonce": 12306, "path": "/v1/datastreams/plug-status/datapoint/", "method": "POST", "body": {"datapoint": {"x": 14}}, "meta": {"Authorization": "token 10994e8b9cfdbb3029741f2223da65c8ddd27338"}}
# {"nonce": 12306, "path": "/v1/datastreams/plug-status/datapoint/", "method": "GET", "meta": {"Authorization": "token 10994e8b9cfdbb3029741f2223da65c8ddd27338"}}

if __name__ == '__main__':
    plug_status = 1
    _sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    _sock.setblocking(1)
    _sock.settimeout(None)
    _sock.connect((HOST, PORT))
    _sock.sendall('{"nonce": 12306, "path": "/v1/device/", "method": "GET", "meta": {"Authorization": "token %s"}}\n' % TOKEN)
    request = ''
    while True:
        _buffer = _sock.recv(4096)
        if not _buffer:
            print 'not _buffer'
            sys.exit(128)
        findIndex = _buffer.find('\n')
        if findIndex < 0:
            request += _buffer
            continue
        request += _buffer[0:findIndex]
        requeststr = request
        request = _buffer[findIndex+1:]
        #print 'raw request: %s' % requeststr
        jsonobj = json.loads(requeststr)
        if 'is_query_device' not in jsonobj:
            #print 'query response from server:\n%s' % (requeststr)
            continue
        print 'get query request from server:\n%s' % (requeststr)
        path = jsonobj['path']
        method = jsonobj['method']
        nonce = jsonobj['nonce']
        m = re.match('^/v1/datastreams/([a-zA-Z0-9_\+\-\.\:]+)/datapoint/?$', path)
        if m:
            stream_name = m.group(1)
            if stream_name == 'plug-status':
                if method == 'GET':
                    print 'get current plug_status %s' % plug_status
                elif method == 'POST':
                    plug_status = jsonobj['body']['datapoint']['x']
                    print 'plug_status change to %s' % plug_status
            response = '{"status": 200, "datapoint": {"x": %s}, "nonce": %s, "is_query_device": true}\n' % (plug_status, nonce)
            _sock.sendall(response)
            #print response

    _sock.sendall('{"status": 400, "result": "failed", "message": "what are you talking?"}\n')
    _sock.close()


