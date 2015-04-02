# -*- coding: utf-8 -*-  
import json
import socket

HOST, PORT = 'ec2-52-74-94-231.ap-southeast-1.compute.amazonaws.com', 80
TOKEN = 'd0f10bf2acc8d8b7ee440fdcc29d0110ccb97ee2'

if __name__ == '__main__':
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.setblocking(1)
    s.settimeout(None)
    print '1. connect to server'
    s.connect((HOST, PORT))
    print '2. identify device with master device key'
    s.sendall('{"router": 10086, "nonce": 12306, "path": "/v1/device/identify/", "method": "GET", "meta": {"Authorization": "token %s"}}\n' % TOKEN)
    print s.recv(4096)
    while True:
        jsonobj = json.loads(s.recv(4096))
        method, path, nonce = jsonobj['method'], jsonobj['path'], jsonobj['nonce']
        if path == '/v1/device/rpc/':
            action = jsonobj['get']['action']
            print 'rpc API with action: %s' % action
        elif path == '/v1/datastreams/multiple/datapoint/':
            datapoint = jsonobj['body']['datapoint']
            print 'datapoint: x: %s, y: %s, z: %s, k: %s, l: %s' % (datapoint['x'], datapoint['y'], datapoint['z'], datapoint['k'], datapoint['l'])
        # ... elif ...
        response = '{"status": 200, "nonce": %s, "deliver_to_device": true}\n' % nonce
        s.sendall(response)
