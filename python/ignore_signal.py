import time
import signal

def sighandler(signum, stack):
    print 'Caught signal %s, ignoring.' %(str(signum))

for i in [x for x in dir(signal) if x.startswith("SIG")]:
    try: 
        signum = getattr(signal, i)
        if signum == 0:
            continue
        signal.signal(signum, sighandler)
    except RuntimeError, m:
        print "Skipping %s" % i

