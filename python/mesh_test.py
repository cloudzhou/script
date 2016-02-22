import struct
import binascii
import socket
import binascii
from iotbucket.mesh.packet import Packet
from iotbucket.mesh.packet import PacketReader

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(('iot.espressif.cn', 7000))
respPacket = Packet(ver=0x1, flags=0x2, proto=0x6, dst_addr='afafafafafaf', src_addr='fafafafafafa', data='{"method": "GET", "path": "/v1/ping"}')
s.sendall(respPacket.dumps())
resp = s.recv(1024)
print resp

reader = PacketReader()
packets = reader.read(resp)
print packets
packet = packets[0]
print packet
print packet.ver
print packet.option
print packet.flags
print packet.proto
print binascii.hexlify(packet.dst_addr)
print binascii.hexlify(packet.src_addr)
print packet.data
