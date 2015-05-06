# -*- coding: utf-8 -*-  
import re
import sys
import ssl
import time
import json
import random
import socket
import httplib
from datetime import datetime

HOST = 'iot.espressif.cn'
PORT = 8000
TOKEN1 = '513d09340e29eb61f91f5cb4e717682c48444d6f'
TOKEN2 = '44590e7ccb254203073aeb61777a4c425b20a018'
MBOX_NAME = 'c08c4b6'

SERVER_PING_URL = '{"nonce": %s, "path": "/v1/ping/", "method": "GET", "meta": {"Authorization": "token %s"}}\n'
IDENTIFY_URL = '{"nonce": %s, "path": "/v1/device/identify/", "method": "POST", "meta": {"Authorization": "token %s"}}\n'
ACK_DELIVER_URL = '{"nonce": %s, "status": 200, "deliver_to_device": true}\n'
MBOX_SUBSCRIBE = '{"path": "/v1/mbox/", "method": "POST", "body": {"name": "%s", "action": "subscribe"}, "meta": {"Authorization": "token %s"}}\n'
MBOX_PUBLISH = '{"path": "/v1/mbox/", "method": "POST", "body": {"name": "%s", "action": "publish", "data": "%s"}, "meta": {"Authorization": "token %s"}}\n'
MBOX_UNSUBSCRIBE = '{"path": "/v1/mbox/", "method": "POST", "body": {"name": "%s", "action": "unsubscribe"}, "meta": {"Authorization": "token %s"}}\n'

class Analoguer():

    def __init__(self, token):
        self.token = token
        self._lines = []
        self._buffer = ''
        self._socket = None

    def quit(self):
        self.send('\x06')

    def read(self):
        while len(self._lines) == 0:
            recv = self._socket.recv(4096)
            assert(recv is not None and recv != '')
            self._buffer += recv
            find_newline = self._buffer.rfind('\n')
            if find_newline == -1:
                continue
            for x in self._buffer[0:find_newline].split('\n'):
                self._lines.append(x)
            self._buffer = self._buffer[find_newline+1:]

        while len(self._lines) > 0:
            pop_line = self._lines.pop()
            if pop_line is None or pop_line == '':
                continue
            return json.loads(pop_line)

        return None

    def send(self, data):
        if not data.endswith('\n'):
            data = data + '\n'
        self._socket.sendall(data)

    def connect(self):
        self._socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self._socket.settimeout(5)
        self._socket.connect((HOST, PORT))

    def close(self):
        if not self._socket:
            return
        try:
            self._socket.close()
        except Exception, e:
            pass
        self._socket = None

    def ping(self):
        randint = self.get_randint()
        self.send(SERVER_PING_URL % (randint, self.token))
        return randint

    def identify(self):
        randint = self.get_randint()
        self.send(IDENTIFY_URL % (randint, self.token))
        return randint

    def subscribe(self, name):
        self.send(MBOX_SUBSCRIBE % (name, self.token))

    def publish(self, name, data):
        self.send(MBOX_PUBLISH % (name, data, self.token))

    def unsubscribe(self, name):
        self.send(MBOX_UNSUBSCRIBE % (name, self.token))

    def ack_deliver(self, nonce):
        self.send(ACK_DELIVER_URL % (nonce))

    def get_randint(self):
        return random.randint(1, 100000000)

    def get_randstr(self, bits_len):
        return '%x' % random.getrandbits(bits_len)

    def get_token(self):
        return ('%x' % random.getrandbits(200))[0:40]


def test_mbox(analoguer1, analoguer2):
    if hasattr(socket, 'setdefaulttimeout'):
        socket.setdefaulttimeout(3)
    total = 0
    success = 0
    ping = 0
    exception_occur = True
    while True:
        total = total + 1
        try:
            socket.gethostbyname('www.baidu.com')
            ping = ping + 1
            if exception_occur:
                analoguer1.close(); analoguer1.connect(); analoguer1.identify(); analoguer1.read(); analoguer1.subscribe(MBOX_NAME); analoguer1.read()
                analoguer2.close(); analoguer2.connect(); analoguer2.identify(); analoguer2.read(); analoguer2.subscribe(MBOX_NAME); analoguer2.read()
            data = analoguer1.get_randstr(30)
            if total % 2 == 0:
                analoguer1.publish(MBOX_NAME, data)
                r = analoguer1.read()
                assert(r['peers'] == 1)
                r = analoguer2.read()
                assert(r['body']['data'] == data)
            else:
                analoguer2.publish(MBOX_NAME, data)
                r = analoguer2.read()
                assert(r['peers'] == 1)
                r = analoguer1.read()
                assert(r['body']['data'] == data)
            exception_occur = False
            success = success + 1
        except Exception, e:
            exception_occur = True
            print e
        print 'total: %s, success: %s, ping: %s, percentage: %s%%' % (total, success, ping, int(success*10000/total)/100.0)
        time.sleep(3)

if __name__ == '__main__':
    analoguer1 = Analoguer(TOKEN1)
    analoguer2 = Analoguer(TOKEN2)
    test_mbox(analoguer1, analoguer2)

