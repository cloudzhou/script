package gracefullyChan

import (
	"sync/atomic"
)

type rawChanWithStoped struct {
	stoped int32
	c      chan interface{}
}

func newRawChanWithStoped() gracefullyChan {
	return &rawChanWithStoped{c: make(chan interface{})}

}
func (r *rawChanWithStoped) start(f func(i interface{})) error {
	go func() {
		for i := range r.c {
			f(i)
		}
	}()
	return nil
}
func (r *rawChanWithStoped) push(i interface{}) bool {
	if atomic.LoadInt32(&r.stoped) == 1 {
		return false
	}
	r.c <- i
	return true
}

func (r *rawChanWithStoped) stop() error {
	atomic.StoreInt32(&r.stoped, 1)
	close(r.c)
	return nil
}
