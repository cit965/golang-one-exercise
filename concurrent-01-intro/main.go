package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个带缓冲的管道，可以缓存2个整数
	ch := make(chan int, 2)

	// 启动一个Goroutine往管道中发送数据
	go func() {
		ch <- 1
		ch <- 2
		fmt.Println("3s后输入4")
		time.Sleep(time.Second * 3)
		ch <- 4
		fmt.Println("输入3完毕")
	}()

	// 从管道中读取数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) // 由于管道中没有数据可读，这个读取操作会导致阻塞
}
