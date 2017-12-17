package gracefullyChan

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// go test -race -timeout 10s -run ^TestRawChanWithStopedAndQuite$

var std gracefullyChan
var duration time.Duration
var new func() gracefullyChan

func testRules(t *testing.T) {
	for i := 0; i < 10; i++ {
		duration = time.Microsecond * time.Duration(i)
		test(t, 0, 0)
		test(t, 0, 1)
		test(t, 1, 0)
		test(t, 1, 1)
		test(t, 0, 10)
		test(t, 10, 0)
		test(t, 10, 10)
		test(t, 1, 10)
		test(t, 10, 1)
		test(t, 0, 1024)
		test(t, 1024, 0)
		test(t, 1024, 1024)
		test(t, 1024, 1024)
		for i := 0; i < 10; i++ {
			chanlen := int(rand.Int31n(10))
			goroutineLen := int(rand.Int31n(10))
			test(t, chanlen, goroutineLen)
		}
	}
}

func test(t *testing.T, chanlen int, goroutineLen int) {
	std = new()
	var result int32
	std.start(func(i interface{}) {
		atomic.AddInt32(&result, 1)
	}, chanlen)
	var total int32
	var wg sync.WaitGroup
	wg.Add(goroutineLen)
	for i := 0; i < goroutineLen; i++ {
		j := i
		go func(k int) {
			pushed := std.push(k)
			if pushed {
				atomic.AddInt32(&total, 1)
			}
			wg.Done()
		}(j)
	}
	time.Sleep(duration)
	std.stop()
	wg.Wait()
	result = atomic.LoadInt32(&result)
	total = atomic.LoadInt32(&total)
	if result != total {
		log.Fatal("result != total ", result, total, chanlen, goroutineLen)
	}
	log.Println("count:", atomic.LoadInt32(&total))
}

func TestRawChan(t *testing.T) {
	new = newRawChan
	testRules(t)
}

func TestRawChanWithStoped(t *testing.T) {
	new = newRawChanWithStoped
	testRules(t)
}

func TestRawChanWithStopedAndWait(t *testing.T) {
	new = newRawChanWithStopedAndWait
	testRules(t)
}

func TestRawChanWithStopedAndLock(t *testing.T) {
	new = newRawChanWithStopedAndLock
	testRules(t)
}

func TestRawChanWithStopedAndQuite(t *testing.T) {
	new = newRawChanWithStopedAndQuite
	testRules(t)
}
