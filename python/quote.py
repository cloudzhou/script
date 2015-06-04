# -*- coding: utf-8 -*-  

import json
import urllib2

sh_codes = ['000001', '601898'] 
sz_codes = ['002030']

class bcolors:
    HEADER = '\033[95m'
    OKBLUE = '\033[94m'
    OKGREEN = '\033[92m'
    OKRED = '\033[31m'
    WARNING = '\033[93m'
    FAIL = '\033[91m'
    ENDC = '\033[0m'
    BOLD = '\033[1m'
    UNDERLINE = '\033[4m'

paramter = ','.join(['%s.sh' % x for x in sh_codes] + ['%s.sz' % x for x in sz_codes])

url = 'http://bdcjhq.hexun.com/quote?s2=%s' % paramter
# print url
c = urllib2.urlopen(url).read().decode('gbk')
keys = ["na", "pc", "op", "vo", "tu", "hi", "lo", "la", "type", "time", "sy", "lt", "sz", "hs", "is"] 
for x in keys:
    c = c.replace('%s:' % x, '"%s":' % x)

l = c.find('({')
r = c.find('})')
c = c[l+1:r+1]
j = json.loads(c)

def normalize(*args):
    n = []
    for x in args:
        x = float(x)
        x = int(x * 100) / 100.0
        n.append(x)
    return tuple(n)

def print_code(j, code):
    if code in j:
        na, la, op, pc, hi, lo = j[code]['na'], j[code]['la'], j[code]['op'], j[code]['pc'], j[code]['hi'], j[code]['lo']
        la, op, pc, hi, lo = normalize(la, op, pc, hi, lo)
        color = bcolors.OKGREEN
        up = la - pc
        up_s = str(up)
        if up > 0:
            up_s = '+' + up_s
            color = bcolors.OKRED
        up_r = int(up * 10000.0 / pc) / 100.0
        up_r_s = '%s%%' % up_r
        if up_r > 0:
            up_r_s = '+' + up_r_s
        print color + '%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s' % (na, la, up_s, up_r_s, op, pc, hi, lo) + bcolors.ENDC

print u'股票\t\t当前\t涨幅\t比率\t今开\t昨收\t最高\t最低'

for code in sh_codes:
    code = code + '.sh'
    print_code(j, code)

for code in sz_codes:
    code = code + '.sz'
    print_code(j, code)

