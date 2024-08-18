package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	val uint
	mu  sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{
		val: 0,
		mu:  sync.Mutex{},
	}
	for range 10000 {
		wg.Add(1)
		go counter.increment(&wg)
	}

	wg.Wait()
	fmt.Println(counter.val)
}

func (c *Counter) increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	c.val++
	c.mu.Unlock()
}
