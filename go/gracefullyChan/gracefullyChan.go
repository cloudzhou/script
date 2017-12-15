package gracefullyChan

type gracefullyChan interface {
	start(f func(i interface{})) error
	push(i interface{}) bool
	stop() error
}
