package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println("goroutine", num, "start")
			// do some work
			fmt.Println("goroutine", num, "done")
		}(i)
	}

	wg.Wait()
	fmt.Println("all goroutines finished")
}
