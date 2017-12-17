package gracefullyChan

import (
	"sync"
)

type rawChanWithStopedAndLock struct {
	stoped int32
	c      chan interface{}
	sync.Mutex
}

func newRawChanWithStopedAndLock() gracefullyChan {
	r := &rawChanWithStopedAndLock{}
	return r
}

func (r *rawChanWithStopedAndLock) start(f func(i interface{}), chanlen int) error {
	r.c = make(chan interface{}, chanlen)
	go func() {
		for i := range r.c {
			f(i)
		}
	}()
	return nil
}

func (r *rawChanWithStopedAndLock) push(i interface{}) bool {
	r.Lock()
	defer r.Unlock()
	if r.stoped == 1 {
		return false
	}
	r.c <- i
	return true
}

func (r *rawChanWithStopedAndLock) stop() error {
	r.Lock()
	defer r.Unlock()
	r.stoped = 1
	close(r.c)
	return nil
}
