#!/usr/bin/python3

import socket
import time

HOST = 'api.vklmotor.com'
PORT = 6000
TOKEN = '0123456789abcde'

def connect(s):
    b = bytearray()
    b.extend([0x43, 0x10, 0x01])
    b.extend(b'0123456789abcde')
    b.extend([0x2b, 0x68])
    s.sendall(b)

def pub(s):
    pass

def ack(s):
    while True:
        buf = s.recv(4096)
        if buf is None:
            return
        cmd = buf[0]
        key = buf[2]
        value = buf[3]
        if cmd == 0x6D:
            if key == 0xE1:
                if value == 0xBA: #lock
                    print('lock')
                elif value == 0xB5: #unlock
                    print('unlock')
            elif key == 0xE2: #beep
                    print('beep')
            elif key == 0xEA: #ping
                    print('ping')
        msgidlen = buf[5]
        msgid = buf[6:6+msgidlen]
        b = bytearray()
        b.extend([0x61, msgidlen+2, 0xE8, msgidlen])
        b.extend(msgid)
        b.extend([0x0, 0x0])
        s.sendall(b)

if __name__ == '__main__':
    while True:
        time.sleep(1)
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.setblocking(1)
        s.settimeout(600000)
        s.connect((HOST, PORT))

        connect(s)
        ack(s)

