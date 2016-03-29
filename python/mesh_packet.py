import struct
import binascii

class Packet(object):

    def __init__(self, *args, **kwargs):
        if 'ver' in kwargs:
            self.ver = kwargs['ver']
        if 'flags' in kwargs:
            self.flags = kwargs['flags']
        if 'proto' in kwargs:
            self.proto = kwargs['proto']
        if 'dst_addr' in kwargs:
            self.dst_addr = kwargs['dst_addr']
        if 'src_addr' in kwargs:
            self.src_addr = kwargs['src_addr']
        self.length = 16
        self.option = 0
        self.ot_len = 0
        if 'option_list' in kwargs and kwargs['option_list'] != None:
            self.option = 1
            self.option_list = kwargs['option_list']
            self.ot_len = len(self.option_list) + 2
            self.length = self.length + self.ot_len
        if 'data' in kwargs:
            self.data = kwargs['data']
            self.length = self.length + len(self.data)

    def dumps(self):
        b = bytearray()
        #header = self.flags | (self.ver << 6) | (self.option << 5)
        header = self.ver | self.option << 2 | self.flags << 3
        b.extend(struct.pack('>B', header))
        b.extend(struct.pack('>B', self.proto))
        b.extend(struct.pack('<H', self.length))
        dst_addr = self.dst_addr
        if len(self.dst_addr) == 12:
            dst_addr = binascii.unhexlify(dst_addr)
        b.extend(dst_addr)
        src_addr = self.src_addr
        if len(self.src_addr) == 12:
            src_addr = binascii.unhexlify(src_addr)
        b.extend(src_addr)
        if self.option:
            b.extend(struct.pack('<H', self.ot_len))
            option_list = self.option_list
            if isinstance(self.option_list, unicode):
                option_list = option_list.decode('utf8')
            b.extend(option_list)
        if self.data:
            data = self.data
            if isinstance(self.data, unicode):
                data = data.decode('utf8')
            b.extend(data)
        return b

class PacketReader():

    def __init__(self):
        self.waiting = PacketEnum.WAITING_HEADER
        self.packet = Packet()
        self.packets = []
        self.tmp = bytearray()
        self.field_len = PacketEnum.FIELD_LEN[self.waiting]

    def read(self, recv_buffer):
        i = 0
        l = len(recv_buffer)
        while i < l:
            byte = recv_buffer[i]
            if self.waiting == PacketEnum.WAITING_HEADER:
                done = self.feed(byte, PacketEnum.WAITING_PROTO, None)
                if done:
                    value = self.field()
                    self.packet.header = value[0]    
                    self.packet.ver = self.packet.header & 0x03
                    self.packet.option = (self.packet.header & 0x4) >> 2
                    self.packet.flags = self.packet.header >> 3
            elif self.waiting == PacketEnum.WAITING_PROTO:
                done = self.feed(byte, PacketEnum.WAITING_LEN, None)
                if done:
                    value = self.field()
                    self.packet.proto = value[0]
            elif self.waiting == PacketEnum.WAITING_LEN:
                done = self.feed(byte, PacketEnum.WAITING_DST_ADDR, None)
                if done:
                    f = self.field()
                    self.packet.length,  = struct.unpack_from('<H', buffer(f))
            elif self.waiting == PacketEnum.WAITING_DST_ADDR:
                done = self.feed(byte, PacketEnum.WAITING_SRC_ADDR, None)
                if done:
                    self.packet.dst_addr = self.field()
            elif self.waiting == PacketEnum.WAITING_SRC_ADDR:
                if self.packet.option == 0:
                    done = self.feed(byte, PacketEnum.WAITING_DATA, self.packet.length - 16)
                    if done:
                        self.packet.src_addr = self.field()
                else:
                    done = self.feed(byte, PacketEnum.WAITING_OT_LEN, None)
                    if done:
                        self.packet.src_addr = self.field()
            elif self.waiting == PacketEnum.WAITING_OT_LEN:
                done = self.feed(byte, PacketEnum.WAITING_OPTION_LIST, None)
                if done:
                    f = self.field()
                    self.packet.ot_len, = struct.unpack_from('<H', buffer(f))
                    self.field_len = self.packet.ot_len - 2
            elif self.waiting == PacketEnum.WAITING_OPTION_LIST:
                done = self.feed(byte, PacketEnum.WAITING_DATA, self.packet.length - 16 - self.packet.ot_len)
                if done:
                    self.packet.option_list = self.field()
            elif self.waiting == PacketEnum.WAITING_DATA:
                done = self.feed(byte, PacketEnum.WAITING_HEADER, None)
                if done:
                    self.packet.data = self.field()
                    self.packets.append(self.packet)
                    self.packet = Packet()
            i += 1

        if len(self.packets) == 0:
            return []
        packets = self.packets
        self.packets = []
        return packets

    def feed(self, byte, next_waiting, field_len):
        self.tmp.append(byte)
        self.field_len -= 1
        if self.field_len == 0:
            if next_waiting is not None:
                self.waiting = next_waiting
            if next_waiting in PacketEnum.FIELD_LEN:
                self.field_len = PacketEnum.FIELD_LEN[next_waiting]
            if field_len is not None:
                self.field_len = field_len
            return True
        return False

    def field(self):
        value = self.tmp
        self.tmp = bytearray()
        return value
        
class PacketEnum:
    WAITING_HEADER = 0x0
    WAITING_PROTO = 0x1
    WAITING_LEN = 0x2
    WAITING_DST_ADDR = 0x3
    WAITING_SRC_ADDR = 0x4
    WAITING_OT_LEN = 0x5
    WAITING_OPTION_LIST = 0x6
    WAITING_DATA = 0x7

    FIELD_LEN = {
        WAITING_HEADER:     1,
        WAITING_PROTO:      1,
        WAITING_LEN:        2,
        WAITING_DST_ADDR:   6,
        WAITING_SRC_ADDR:   6,
        WAITING_OT_LEN:     2,
    }
