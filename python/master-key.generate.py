import os
import json
import binascii
import string
import httplib
import urllib2
from datetime import datetime

USER_TOKEN = 'your user token' # user token
PRODUCT_ID = 0 # product id you want to generate device
devices_count = 1 # count

DATETIME_FORMAT = '%Y%m%d-%H-%M-%S'
now = datetime.now()
dir_datetime = now.strftime(DATETIME_FORMAT)

abspath = os.path.abspath(__file__)
dname = os.path.dirname(abspath) + os.sep + dir_datetime
os.makedirs(dname)
os.chdir(dname)

host = 'iot.espressif.cn'
headers = {'Host': host, 'User-Agent': 'Mozilla', 'Accept': '*/*', 'Authorization': 'token ' + USER_TOKEN}
devices_serial = ['{}'] * devices_count
params = '{"devices": ['+ ','.join(devices_serial) + ']}'
path = '/v1/products/%s/devices/' % PRODUCT_ID
conn = httplib.HTTPConnection(host, port=80)
conn.request("POST", path, params, headers)
response = conn.getresponse()
data = response.read()
devices_as_dict = json.loads(data)
devices = []
print data
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

devices_generate = open('devices-generate.csv', 'w')
devices_generate.write('id,serial,token\n')
for device in devices:
    devices_generate.write('%s,%s,%s\n' % (device['id'], device['serial'], device['token']))
devices_generate.close()


