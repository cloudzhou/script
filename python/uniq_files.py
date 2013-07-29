# -*- coding: utf-8 -*-  
#!/usr/bin/python

import os, sys, codecs, re

basepath = '/tmp/ss'
uniq_line = {}
for filepath in os.listdir(basepath):
    abs_filepath = basepath + '/' + filepath
    if not os.path.isfile(abs_filepath) or not re.match('.+\.py$', abs_filepath):
        continue
    filtered_line = []
    with codecs.open(abs_filepath, encoding='utf-8') as readF:
        for line in readF:
            if line in uniq_line:
                continue
            uniq_line[line] = 1
            filtered_line.append(line)
    with codecs.open(abs_filepath, encoding='utf-8', mode='w+') as writeF:
        writeF.write(''.join(filtered_line))
        writeF.flush()
    print 'aaa'
    print 'bbb'
    print 'ccc'
