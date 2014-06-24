import socket, ssl, pprint

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
ssl_sock = ssl.wrap_socket(s,
                           ca_certs='/home/cloudzhou/script/python/ssl_demo/suibian.com.crt',
                           cert_reqs=ssl.CERT_REQUIRED)
ssl_sock.connect(('localhost', 9443))
print repr(ssl_sock.getpeername())
print ssl_sock.cipher()
print pprint.pformat(ssl_sock.getpeercert())
ssl_sock.write('hello')
data = ssl_sock.read()
print data
ssl_sock.close()


