package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	c := Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Counter value:", c.Value())
}
