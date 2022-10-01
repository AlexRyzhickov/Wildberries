package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	var flag int32 = 0
	waitEvent := func(c <-chan interface{}) {
		v := <-c
		if atomic.CompareAndSwapInt32(&flag, 0, 1) {
			out <- v
			close(out)
		}
	}
	for _, c := range channels {
		go waitEvent(c)
	}
	return out
}

func main() {
	var sig = func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	// Ожидаем запись в канал, что основанная горутина не завершилась всех остальных
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))
}
