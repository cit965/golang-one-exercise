package main

import "fmt"

func main() {
	age := 12
	switch age {
	case 0:
		fmt.Println(0)
	case 28:
		fmt.Println("28了")
		//fallthrough
	}
}
