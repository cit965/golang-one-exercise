package main

import "fmt"

func main() {
	var yourName = "NanGe"
	fmt.Println("your name is :", yourName)

	var myName string = "DaMing"
	fmt.Println("my name is :", myName)

	adminName := "admin"
	fmt.Println("admin name is :", adminName)

	var num int
	fmt.Println("the num is :", num)

	part1, other := 1, "other"
	fmt.Printf("the part1 :%v,other:%v", part1, other)

	var (
		part2 = "part2"
		part3 = "part3"
	)
	fmt.Println(part3, part2)

	par4, part5, _ := "ddd", "df", "dd"
	fmt.Println(par4, part5)
}
