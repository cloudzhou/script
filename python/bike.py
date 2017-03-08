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
    b = bytearray()
    b.extend([0x70, 0x15])
    b.extend([0x71])
    b.extend(b'313437935')
    b.extend([0x72])
    b.extend(b'1202510124')
    b.extend([0x2b, 0x68])
    try:
        s.sendall(b)
    except: 
        return True
    return False

def ack(s):
    while True:
        try:
            buf = s.recv(4096)
        except: 
            if pub(s):
                return
            continue
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
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.setblocking(1)
        s.settimeout(6)
        s.connect((HOST, PORT))

        connect(s)
        ack(s)
        time.sleep(1)
