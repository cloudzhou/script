import time
import struct
import paho.mqtt.client as mqtt
from paho.mqtt.client import MQTTv311

KEY = 'e7e033b3e5c90a313e81'
HOST = 'iot.espressif.cn'
MBOX = 'mbox/bb57677'

def on_connect(client, userdata, flags, rc):
    print('on_connect result code: %s' % str(rc))
    # client.subscribe(MBOX, qos=1)

def on_disconnect(client, userdata, rc):
    print('on_disconnect result code: %s' % str(rc))

def on_message(client, userdata, message):
    print('on_message message: %s' % message)
    print message.topic
    if message.topic.startswith('$datapoint/'):
        l = len(message.payload) / 4
        print struct.unpack('!' + 'i' * l, message.payload) 
    if message.topic == '$rpc':
        print message.payload
    if message.topic.startswith('mbox/'):
        print message.payload

def on_publish(client, userdata, mid):
    print('on_publish mid: %s' % mid)

def on_subscribe(client, userdata, mid, granted_qos):
    print('on_subscribe mid: %s, granted_qos: %s' % (mid, granted_qos))

def on_unsubscribe(client, userdata, mid):
    print('on_unsubscribe mid: %s' % mid)

client = mqtt.Client(client_id=KEY, protocol=MQTTv311)
client.on_connect = on_connect
client.on_disconnect = on_disconnect
client.on_message = on_message
client.on_publish = on_publish
client.on_subscribe = on_subscribe
client.on_unsubscribe = on_unsubscribe

client.connect(HOST, 1883, 10)
i = 0
while i < 100:
    client.loop()
    client.publish(MBOX, payload='hello' + str(i), qos=1, retain=False)
    i = i + 1
    time.sleep(1)

# datapoint publish
"""
    payload = struct.pack('!iii', 1, 2, 3)
    client.publish('$datapoint/mulit-aa-bb', payload=payload, qos=1, retain=False)
"""

