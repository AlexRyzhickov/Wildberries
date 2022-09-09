package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count uint64
}

func (c *Counter) Increment() {
	atomic.AddUint64(&c.count, 1)
}

func (c *Counter) Count() uint64 {
	return atomic.LoadUint64(&c.count)
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{count: 0}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				counter.Increment()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter.Count())
}
