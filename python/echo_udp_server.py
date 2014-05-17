"""A simple UDP server.

For every message received, it sends a reply back.

You can use udp_client.py to send a message.
"""
from __future__ import print_function
from gevent.server import DatagramServer


class EchoServer(DatagramServer):

    def handle(self, data, address):
        if not data:
            print('not data')
            return
        print('%s: got %r' % (address[0], data))
        self.socket.sendto(data, address)


if __name__ == '__main__':
    print('Receiving datagrams on :9000')
    EchoServer('0.0.0.0:9003').serve_forever()
