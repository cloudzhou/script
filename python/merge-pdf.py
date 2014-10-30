# -*- coding: utf8 -*-
import os
import re
import sys
import codecs

ACROBAT_EXE = 'C:\\Program Files (x86)\\Adobe\\Acrobat 11.0\\Acrobat\\Acrobat.exe'
GXPS_EXE = 'C:\\Program Files\\gs\\gs9.15\\bin\\gxps-9.15-win32.exe'


SOURCE_PATH = u'D:\\new'.replace('\\', '/')
XPS_PATH = u'D:\\xps'

PDF2XPS_BAT = u'%s/%s' % (SOURCE_PATH, 'pdf2xps.bat')
XPS2PDF_BAT = u'%s/%s' % (SOURCE_PATH, 'xps2pdf.bat')
XPS_PDF_RULE = u'%s/%s' % (SOURCE_PATH, 'xps2pdf.rule.txt')

os.chdir(SOURCE_PATH)

# pdf to xps #
def render_pdf2xps_bat():
    count = 0
    pdf2xps_bat_w = codecs.open(PDF2XPS_BAT, encoding='gb18030', mode='w')
    for path, dirs, files in os.walk(SOURCE_PATH):
        for file in files:
            if not file.endswith('.pdf'):
                continue
            abs_path = path + '\\' + file
            abs_path = abs_path.replace('/', '\\')        
            count = count + 1
            xps_file = u'%s\\%s.xps' % (XPS_PATH, count)
            abs_path = abs_path.replace('/', '\\')
            xps_file = xps_file.replace('/', '\\')
            pdf2xps_bat_w.write('"%s" /n /t "%s" "Microsoft XPS Document Writer" "Microsoft XPS Document Writer" "%s"\r\n' % (ACROBAT_EXE, abs_path, xps_file))
    pdf2xps_bat_w.close()

# merge pdf #
def get_index_slash(string, slash_count):
    i = 0
    count = 0
    while i < len(string):
        if string[i] == '\\':
            count = count + 1
        if count == slash_count:
            return i
        i = i + 1

def get_last_slash(string):
    i = len(string) - 1
    while i >= 0:
        if string[i] == '\\':
            return i
        i = i - 1
    return len(string)

def fillwith_f_d():
    f = {}
    d = {}

    pdf2xps_bat_r = codecs.open(PDF2XPS_BAT, encoding='gb18030', mode='r')
    for x in pdf2xps_bat_r:
        x = x.strip()
        xx = x.split('"')
        pdf_path = xx[3]
        xps_path = xx[9]
        count = pdf_path.count('\\')
        #if count >= 6:
        #    count = 6
        index = get_index_slash(pdf_path, count)
        base_pdf_path = pdf_path[0:index]
        pdf_file = pdf_path[index:]
        if base_pdf_path not in d:
            d[base_pdf_path] = []
        f[pdf_path] = xps_path
        d[base_pdf_path].append(pdf_file)
    pdf2xps_bat_r.close()

    dd = {}

    for x in d:
        files = d[x]
        without_num = []
        contain_num = {}
        for y in files:
            m = re.match('.*?(\d+)_*\.pdf', y)
            if m:
                num = int(m.group(1))
                if num not in contain_num:
                    contain_num[num] = []
                contain_num[num].append(y)
            else:
                without_num.append(y)
        for xx in contain_num:
            contain_num[xx].sort()

        ordered_files = []
        nums = contain_num.keys()
        nums.sort()
        for num in nums:
            ordered_files += contain_num[num]

        without_num.sort()
        ordered_files += without_num

        dd[x] = ordered_files

    d = dd
    return f, d

def render_xps2pdf_rule(f, d):
    xps2pdf_rule_w = codecs.open(XPS_PDF_RULE, encoding='gb18030', mode='w')
    for x in d:
        pdfs = []
        xps = []
        for y in d[x]:
            pdf_file_path = x + y
            pdfs.append(pdf_file_path)
            xps.append(f[pdf_file_path])
        new_dir = x[0:2] + '\\tmp' + x[2:]
        index = get_last_slash(new_dir)
        pdf_file = new_dir + '.pdf'
        xps2pdf_rule_w.write('---- %s ----\r\n' % pdf_file)
        for y in pdfs:
            xps2pdf_rule_w.write('%s __|;|__ %s\r\n' % (y, f[y]))
        xps2pdf_rule_w.write('\r\n')
    xps2pdf_rule_w.close()

def write_gxps_script(xps2pdf_bat_w, target_pdf, target_xps):
    if target_pdf is None:
        return
    pdf_file = target_pdf[0:2] + '\\tmp' + target_pdf[2:]
    index = get_last_slash(pdf_file)
    parent_dir = pdf_file[0:index]
    xps_files = ' '.join(target_xps)
    xps2pdf_bat_w.write('mkdir %s\r\n' % parent_dir)
    xps2pdf_bat_w.write('"%s" -o "%s" -sDEVICE=pdfwrite %s\r\n' % (GXPS_EXE, pdf_file, xps_files))

def render_xps2pdf_bat():
    xps2pdf_rule_r = codecs.open(XPS_PDF_RULE, encoding='gb18030', mode='r')
    xps2pdf_bat_w = codecs.open(XPS2PDF_BAT, encoding='gb18030', mode='w')

    target_pdf = None
    target_xps = []
    for x in xps2pdf_rule_r:
        x = x.strip()
        if x == '':
            continue
        if x.startswith('---- ') and x.endswith(' ----'):
            write_gxps_script(xps2pdf_bat_w, target_pdf, target_xps)
            target_pdf = x[5:len(x)-5]
            target_xps = []
            continue
        if ' __|;|__ ' in x:
            (pdf, xps) = x.split(' __|;|__ ')
            target_xps.append(xps)

    write_gxps_script(xps2pdf_bat_w, target_pdf, target_xps)

    xps2pdf_bat_w.close()
    xps2pdf_rule_r.close()

if __name__ == '__main__':
    # 1 render pdf2xps.bat
    render_pdf2xps_bat()
    # 2 render xps2pdf.rule.txt
    (f, d) = fillwith_f_d()
    render_xps2pdf_rule(f, d)
    # 3 render xps2pdf.bat
    render_xps2pdf_bat()

