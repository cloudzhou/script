# -*- coding: utf-8 -*-  
import codecs
import urllib2
import json
from sets import Set
"""
一个用来检测字符表里面是否有可用的拼音域名的脚本
http://pypi.python.org/pypi/chardet 用来检测编码
"""
f = codecs.open('data/pinyin.dic', encoding='utf-8')
pinyin_dict = {}
for line in f:
    line = line.strip()
    pinyin = line.split(' ')
    for x in pinyin[1:]:
       pinyin_dict[x] = pinyin[0]

domain_pinyins = Set()
f = codecs.open('data/words.dic', encoding='utf-8')
for line in f:
    line = line.strip()
    if len(line) == 2:
        if line[0] in pinyin_dict and line[1] in pinyin_dict:
            domain_pinyin = pinyin_dict[line[0]] + pinyin_dict[line[1]]
            domain_pinyins.add(domain_pinyin)

domain_pinyins = sorted(domain_pinyins)
i = 0
result = open('data/result', 'w+')
for domain_pinyin in domain_pinyins:
    try:
        req = urllib2.Request(url='http://who.is/domains/buying-options.json?terms=' + domain_pinyin + '.com')
        urlf = urllib2.urlopen(req, timeout=10)
        json_content = urlf.read()
        jsono = json.loads(json_content)
        text = jsono[u'available_for_purchase'][u'text']
        print '%s %s' % (i, domain_pinyin)
        if text == 'Available':
            result.write(domain_pinyin)
            result.write('\n')
            result.flush()
    except Exception, e:
        pass
    
