import json
import struct
import binascii
import socket
import binascii
import time
from iotbucket.mesh.packet import Packet
from iotbucket.mesh.packet import PacketReader

host = 'localhost'
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((host, 7000))
reader = PacketReader()

def sendPacket(data):
    p = Packet(ver=0x1, flags=0x2, proto=0x6, dst_addr='a0a0a0a0a0a0', src_addr='0a0a0a0a0a0a', data=data)
    s.sendall(p.dumps())

def readPacket():
    packets = []
    while len(packets) == 0:
        resp = s.recv(1024)
        if not resp:
            return None
        """print 'start --------------------'
        print len(resp)
        print resp
        print 'end --------------------'"""
        packets = reader.read(resp)
    packet = packets[0]
    return packet

def dumps(packet):
    print packet
    print packet.ver
    print packet.option
    print packet.flags
    print packet.proto
    print binascii.hexlify(packet.dst_addr)
    print binascii.hexlify(packet.src_addr)
    if packet.option == 1:
        print packet.option_list
    print packet.data

def deliver_data(packet):
    j = json.loads(str(packet.data))
    nonce = j['nonce']
    data = '{"status": 200,"nonce": %s, "deliver_to_device":true,"mdev_mac":"5CCF7F0A13FE"}\n' % nonce
    return data

sendPacket('{"method": "GET", "path": "/v1/ping", "meta": {"Authorization": "token 7c42a0d69c38422427732b13bcf74845c918ed10"}}')
p = readPacket()
dumps(p)

while True:
    sendPacket('{"method": "GET", "path": "/v1/ping", "meta": {"Authorization": "token 7c42a0d69c38422427732b13bcf74845c918ed10"}}')
    p = readPacket()
    dumps(p)
    time.sleep(1)

while False:
    p = readPacket()
    dumps(p)
    data = deliver_data(p)
    print data
    sendPacket(data)
