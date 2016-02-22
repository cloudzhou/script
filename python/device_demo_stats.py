import sys
import time
import json
import httplib

HOST = 'iot.espressif.cn'
USER_AGENT = 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.80 Safari/537.36'
ACCEPT = 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8'

PATH = '/v1/datastreams/plug-status/datapoint/?deliver_to_device=true'
DATASTREAM = 'plug-status'
OWNER_KEY = 'd3b2f1d0e308c42515adb328973c56c16dbc5872'
ON = {"datapoint":{"x":1}}
OFF = {"datapoint":{"x":0}}
GET, POST = 'GET', 'POST'
LOOP = 100
PERIOD = 3

global total_count, failed_count, on_failed, off_failed, query_failed
total_count, failed_count, on_failed, off_failed, query_failed = 0, 0, 0, 0, 0

def get_rest_resp(method, path, params, token=None):
    if token is None:
        token = ''
    if isinstance(params, dict):
        params = json.dumps(params)
    headers = {'Host': HOST, 'User-Agent': USER_AGENT, 'Accept': ACCEPT, 'Authorization': 'token %s' % token}
    conn = httplib.HTTPConnection(HOST, port=80)
    try:
        conn.request(method, path, params, headers)
        response = conn.getresponse()
        data = response.read()
        return json.loads(data), True
    finally:
        conn.close()
    return None, False

def on():
    resp, r = get_rest_resp(POST, PATH, ON, OWNER_KEY)
    if not r:
        return False
    return resp['status'] == 200

def off():
    resp, r = get_rest_resp(POST, PATH, OFF, OWNER_KEY)
    if not r:
        return False
    return resp['status'] == 200

def query():
    resp, r = get_rest_resp(GET, PATH, '', OWNER_KEY)
    if not r or resp['status'] != 200:
        return -1
    return resp['datapoint']['x']

def test():
    global total_count, failed_count, on_failed, off_failed, query_failed
    if not on():
        on_failed += 1
        return False
    time.sleep(1)
    status = query()
    if status != 1:
        print 'status != 1'
        query_failed += 1
        return False
    if not off():
        off_failed += 1
        return False
    status = query()
    if status != 0:
        query_failed += 1
        return False
    return True

l = 0
while l < LOOP:
    time.sleep(PERIOD)
    total_count += 1
    r = test()
    if not r:
        failed_count += 1
    percent = int((1-failed_count*1.0/total_count)*10000) / 100.0
    log = "\rtotal_count: %s, failed_count:%s, percent: %s%%, on_failed: %s, off_failed: %s, query_failed: %s" % (total_count, failed_count, percent, on_failed, off_failed, query_failed)
    sys.stdout.write(log)
    sys.stdout.flush()
    l += 1
