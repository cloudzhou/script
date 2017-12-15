package gracefullyChan

type rawChan struct {
	c chan interface{}
}

func newRawChan() gracefullyChan {
	return &rawChan{c: make(chan interface{})}
}

func (r *rawChan) start(f func(i interface{})) error {
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
