#!/usr/bin/python
#-*- coding:utf-8 -*-

import re

s = "<a class=f14b href=http://mp.ppsj.com.cn/jsbtfz.html   target=_blank>长江纺织</a>"
r = re.compile("<a class=f14b href=(.+?)\s+target=_blank>(.+?)</a>")
m = r.match(s)
# print(m.groups(0))

f = open("0.html")
content = f.read()
# print(content)

for x in re.findall(r, content):
    print(x[0])
    print(x[1])

y = 0
for x in xrange(0, 189):
    f = open(str(x)+".html")
    content = f.read()
    f.close()
    for xx in re.findall(r, content):
        record = "%s;%s;%s_%s.html" % (xx[1], xx[0], x, y)
        y = y + 1
        print(record)
