from datetime import datetime
import os
import binascii
import string
import httplib
import urllib2
import json

DATETIME_FORMAT = '%Y%m%d-%H-%M-%S'
now = datetime.now()
dir_datetime = now.strftime(DATETIME_FORMAT)

abspath = os.path.abspath(__file__)
dname = os.path.dirname(abspath) + os.sep + dir_datetime
os.makedirs(dname)
os.chdir(dname)


devices_count = 1
host = '114.215.177.97'
headers = {'Host': host, 'User-Agent': 'Mozilla', 'Accept': '*/*', 'Authorization': 'token '}
devices_serial = ['{"serial": "1"}']*devices_count
params = '{"devices": ['+ ','.join(devices_serial) + ']}'
path = '/v1/products/1/devices/'
conn = httplib.HTTPConnection(host, port=80)
conn.request("POST", path, params, headers)
response = conn.getresponse()
data = response.read()
devices_as_dict = json.loads(data)
devices = []
for device in devices_as_dict['devices']:
    (id, serial, token) = (device['id'], device['serial'], device['key']['token'])
    devices.append({'id':str(id), 'serial': str(serial), 'token': str(token)})

bitlist = ['FF']*88
bytes = binascii.a2b_hex(''.join(bitlist))
for device in devices:
    bitout = open(device['token']+'.bin', 'wb')
    bitout.write(device['token'])
    bitout.write(bytes)
    bitout.close()

devices_generate = open('devices-generate.txt', 'w')
for device in devices:
    devices_generate.write(device['id'] + ';' + device['serial'] + ';' + device['token'] + '\n')
devices_generate.close()

