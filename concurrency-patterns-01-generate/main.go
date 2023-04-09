package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func main() {
	out := gen(1, 2, 3, 4, 5)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	// 管道被关闭，打印零值
	fmt.Println(<-out)
	fmt.Println(<-out)
}
