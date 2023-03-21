package main

import "fmt"

func main() {
	// 1写一个函数，能够接受一个人的名字，并打印在显示器中
	print("nan ge")
	// 2写一个函数，能够计算两个数之和
	fmt.Println(sum(12, 13))
	// 3写一个函数，能够计算两个数各自的平方，并返回两个结果
	m, n := pingfang(3, 4)
	fmt.Println(m, n)
	// 4合并使用23
	fmt.Println(sum(pingfang(3, 4)))
}

func pingfang(i, j int) (int, int) {
	return i * i, j * j
}

func sum(i int, i2 int) int {
	return i + i2
}

func print(name string) {
	fmt.Println(name)
}
