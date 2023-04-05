package main

import "fmt"

// 要求实现一个函数，将一个切片中的元素循环右移k个位置
func rotate(s []int, k int) {
	n := len(s)
	if k %= n; k != 0 {
		reverse(s[:n-k])
		reverse(s[n-k:])
		reverse(s)
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("before: %v\n", s)
	rotate(s, 2)
	fmt.Printf("after: %v\n", s)
}

//在这个示例中，我们定义了一个rotate函数，接受一个整型切片和一个整数k作为参数，将切片中的元素循环右移k个位置。具体实现是先将切片分成两部分，分别逆序，然后将整个切片逆序即可。
//在实现逆序的reverse函数中，我们使用了一个简单的双指针算法。
//在main函数中，我们创建了一个长度为5的整型切片s，并将其元素分别初始化为1, 2, 3, 4, 5。
//然后，我们调用rotate函数将切片右移2个位置，并输出移位后的结果。运行上述代码，输出结果如下
//before: [1 2 3 4 5]
//after: [4 5 1 2 3]
