package gracefullyChan

import (
	"fmt"
	"log"
	"sync/atomic"
	"testing"
)

var std gracefullyChan

func test(t *testing.T) {
	std.start(func(i interface{}) {
		log.Println(fmt.Sprintf("%v", i))
	})
	var total int32
	for i := 0; i < 1024; i++ {
		j := i
		go func(i int) {
			pushed := std.push(i)
			if pushed {
				atomic.AddInt32(&total, 1)
			}
		}(j)
	}
	std.stop()
	log.Println("count:", atomic.LoadInt32(&total))
}

func TestRawChan(t *testing.T) {
	std = newRawChan()
	test(t)
}

func TestRawChanWithStoped(t *testing.T) {
	std = newRawChanWithStoped()
	test(t)
}

func TestRawChanWithStopedAndWait(t *testing.T) {
	std = newRawChanWithStopedAndWait()
	test(t)
}

func TestRawChanWithStopedAndLock(t *testing.T) {
	std = newRawChanWithStopedAndLock()
	test(t)
}

func TestRawChanWithStopedAndQuite(t *testing.T) {
	std = newRawChanWithStopedAndQuite()
	test(t)
}
