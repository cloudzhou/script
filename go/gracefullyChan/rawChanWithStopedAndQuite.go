package gracefullyChan

import (
	"sync"
	"sync/atomic"
)

type rawChanWithStopedAndQuite struct {
	counter int32
	stoped  int32
	c       chan interface{}
	quite   interface{}
	wg      *sync.WaitGroup
}

func newRawChanWithStopedAndQuite() gracefullyChan {
	var wg sync.WaitGroup
	wg.Add(1)
	r := &rawChanWithStopedAndQuite{quite: &struct{}{}, wg: &wg}
	return r
}

func (r *rawChanWithStopedAndQuite) start(f func(i interface{}), chanlen int) error {
	r.c = make(chan interface{}, chanlen)
	go func() {
		for i := range r.c {
			if i == r.quite {
				break
			}
			f(i)
		}
		for {
			select {
			case i := <-r.c:
				f(i)
			default:
				if atomic.LoadInt32(&r.counter) == 0 {
					close(r.c)
					r.wg.Done()
					return
				}
			}
		}
	}()
	return nil
}

func (r *rawChanWithStopedAndQuite) push(i interface{}) bool {
	atomic.AddInt32(&r.counter, 1)
	defer atomic.AddInt32(&r.counter, -1)
	if atomic.LoadInt32(&r.stoped) == 1 {
		return false
	}
	r.c <- i
	return true
}

func (r *rawChanWithStopedAndQuite) stop() error {
	atomic.StoreInt32(&r.stoped, 1)
	r.c <- r.quite
	r.wg.Wait()
	return nil
}
