package gracefullyChan

type rawChan struct {
	c chan interface{}
}

func newRawChan() gracefullyChan {
	return &rawChan{}
}

func (r *rawChan) start(f func(i interface{}), chanlen int) error {
	r.c = make(chan interface{}, chanlen)
	go func() {
		for i := range r.c {
			f(i)
		}
	}()
	return nil
}

func (r *rawChan) push(i interface{}) bool {
	r.c <- i
	return true
}

func (r *rawChan) stop() error {
	close(r.c)
	return nil
}
