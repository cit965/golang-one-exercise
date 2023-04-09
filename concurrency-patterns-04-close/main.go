package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 我们知道 发送方 发送完数据可以选择主动关闭，但是如果有的时候需要 接收到某个通知再关闭通道的话 改如何实现呢？
	// 我们使用context

	c, cancel := context.WithCancel(context.Background())
	out := gen(c)

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("start to con")
		cancel()
	}()

	for v := range out {
		fmt.Println(v)
	}

}

func gen(ctx context.Context) chan string {
	out := make(chan string)
	tick := time.Tick(time.Second)
	go func() {
		for {
			select {
			case <-tick:
				out <- "cc"
			case <-ctx.Done():
				out <- "close"
				close(out)
				return
			}
		}
	}()
	return out
}
