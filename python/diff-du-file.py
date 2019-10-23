import subprocess
import time

logdu = {}
last_logdu = {}

def load_logdu():
    out = subprocess.Popen(["/bin/bash", "/tmp/t.sh"], stdout=subprocess.PIPE)
    stdout, stderr = out.communicate()
    for x in stdout.split('\n'):
        y = x.split('\t', 2)
        if len(y) != 2:
            continue
        logdu[y[1]] = int(y[0])


def diff_logdu():
    print '--------------------------------'
    for log, du in logdu.items():
        if log not in last_logdu:
            print '+%d\t%s' % (du, log)
        else:
            last_du = last_logdu[log]
            if du != last_du:
                fmt = '+%d\t+%0.0002f%%\t%s'
                if du < last_du:
                    fmt = '%d\t%0.0002f%%\t%s'
                print fmt % (du-last_du, (du-last_du)*100.0/last_du*1.0, log)
    for log, du in last_logdu.items():
        if log not in logdu:
            print '-%d\t%s' % (du, log)

# load first time
load_logdu()
last_logdu = logdu
logdu = {}
time.sleep(10)

# diff loop 
while True:
    load_logdu()
    diff_logdu()
    last_logdu = logdu
    logdu = {}
    time.sleep(10)

