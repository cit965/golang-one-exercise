package main

import "fmt"

func main() {
	sum := 10
	for sum < 20 {
		sum++
		if sum == 16 {
			continue
		}
		fmt.Println(sum)
	}
}
