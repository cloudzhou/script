#!/usr/bin/env python

import sys, traceback, json
from gevent.server import StreamServer

def echo(socket, address):
    try:
        while True:
            _buffer = socket.recv(4096)
            if not _buffer:
                return
            socket.sendall(_buffer)
    except Exception, e:
        print traceback.format_exc()
    finally:
        socket.close()


if __name__ == '__main__':
    server = StreamServer(('0.0.0.0', 9001), echo)
    server.serve_forever()

