#!/usr/bin/python
import glob, sys, os
from subprocess import check_output, Popen, PIPE

STATIC_PATH = '/opt/app/static/static'
JS_COMPILER = '/opt/run/closure-library/compiler/compiler.jar'
JAVA = 'java'
js = {
    'lib': [
        'js/jquery-1.7.2.min.js',
        'js/bootstrap.min.js',
        'js/select2.min.js',
        'js/app/underscore/underscore-min.js',
        'js/app/underscore/underscore.string.min.js',
        'js/moment.min.js',
        'js/moment.zh-cn.js',
        'js/md5.min.js',
        'js/less-1.4.0.min.js',
    ],
    'markdown': [
        'js/Markdown.Converter.js',
        'js/Markdown.Editor.js',
        'js/Markdown.Sanitizer.js',
    ],
    'highcharts': [
        'js/app/highcharts/highcharts.js',
        'js/app/highcharts/themes/gitshell.js',
    ],
    'syntaxhighlighter': [
        'js/app/syntaxhighlighter/shCore.js',
        'js/app/syntaxhighlighter/shAutoloader.js',
        'js/app/syntaxhighlighter/shLegacy.js',
        'js/app/syntaxhighlighter/shBrush*.js',
    ],
    'gitshell': [
        'js/gitshell-template.js',
        'js/gitshell.js',
    ],
}
LESSC = '/usr/local/bin/lessc'
css = {
    'syntaxhighlighter': [
        'less/app/syntaxhighlighter/shCore.css',
        'less/app/syntaxhighlighter/shThemeDefault.css',
    ],
    'gitshell': [
        'less/gitshell.less',
    ]
}

for key, value in js.items():
    filepath = '%s/js/min/%s.min.js' % (STATIC_PATH, key)
    all_jsfile_paths = []
    for item in value:
        abs_jsfile_paths = glob.glob('%s/%s' % (STATIC_PATH, item))
        for abs_jsfile_path in abs_jsfile_paths:
            if not os.path.exists(abs_jsfile_path):
                print 'filepath: %s is not exists, >&<' % abs_jsfile_path
        all_jsfile_paths = all_jsfile_paths + abs_jsfile_paths        
    command = '%s -jar %s --language_in=ECMASCRIPT5 --js %s --js_output_file %s' % (JAVA, JS_COMPILER, ' '.join(all_jsfile_paths), filepath)
    print command
    check_output(command, shell=True)

for key, value in css.items():
    filepath = '%s/less/min/%s.min.css' % (STATIC_PATH, key)
    all_lessfile_paths = []
    for item in value:
        abs_lessfile_paths = glob.glob('%s/%s' % (STATIC_PATH, item))
        for abs_lessfile_path in abs_lessfile_paths:
            if not os.path.exists(abs_lessfile_path):
                print 'filepath: %s is not exists, >&<' % abs_lessfile_path
        all_lessfile_paths = all_lessfile_paths + abs_lessfile_paths        
    command = '> %s ; for i in %s; do lessc --yui-compress $i >> %s; done' % (filepath, ' '.join(all_lessfile_paths), filepath)
    print command
    check_output(command, shell=True)

