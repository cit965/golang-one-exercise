package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 = 0 // 使用 int64 类型变量作为计数器
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10000; j++ {
				atomic.AddInt64(&counter, 1) // 原子地将计数器加 1
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", atomic.LoadInt64(&counter))
}
