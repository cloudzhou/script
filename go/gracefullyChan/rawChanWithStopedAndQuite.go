package gracefullyChan

import (
	"sync"
	"sync/atomic"
)

type rawChanWithStopedAndQuite struct {
	stoped int32
	c      chan interface{}
	quite  interface{}
	wg     *sync.WaitGroup
}

func newRawChanWithStopedAndQuite() gracefullyChan {
	var wg sync.WaitGroup
	wg.Add(1)
	r := &rawChanWithStopedAndQuite{c: make(chan interface{}, 10), quite: &struct{}{}, wg: &wg}
	return r
}

func (r *rawChanWithStopedAndQuite) start(f func(i interface{})) error {
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
				close(r.c)
				r.wg.Done()
				return
			}
		}
	}()
	return nil
}

func (r *rawChanWithStopedAndQuite) push(i interface{}) bool {
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
