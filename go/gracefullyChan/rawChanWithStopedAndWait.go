package gracefullyChan

import (
	"sync"
	"sync/atomic"
)

type rawChanWithStopedAndWait struct {
	stoped int32
	c      chan interface{}
	wg     sync.WaitGroup
}

func newRawChanWithStopedAndWait() gracefullyChan {
	r := &rawChanWithStopedAndWait{}
	r.wg = sync.WaitGroup{}
	r.wg.Add(1)
	return r
}

func (r *rawChanWithStopedAndWait) start(f func(i interface{}), chanlen int) error {
	r.c = make(chan interface{}, chanlen)
	go func() {
		for i := range r.c {
			f(i)
		}
	}()
	return nil
}
func (r *rawChanWithStopedAndWait) push(i interface{}) bool {
	if atomic.LoadInt32(&r.stoped) > 0 {
		if atomic.CompareAndSwapInt32(&r.stoped, 1, 2) {
			r.wg.Done()
		}
		return false
	}
	r.c <- i
	return true
}

func (r *rawChanWithStopedAndWait) stop() error {
	atomic.StoreInt32(&r.stoped, 1)
	r.wg.Wait()
	close(r.c)
	return nil
}
