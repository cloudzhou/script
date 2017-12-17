package gracefullyChan

type gracefullyChan interface {
	start(f func(i interface{}), chanlen int) error
	push(i interface{}) bool
	stop() error
}
