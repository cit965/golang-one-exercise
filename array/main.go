package main

import "fmt"

type Person struct {
	Name  string
	Money int
}

func main() {
	//1.初始化4个人
	//2.打印下每个人的money,以及总的money
	//3.改变下3，4 人的money
	//4.再次打印下
	persions := [...]Person{
		{
			Name:  "yi",
			Money: 12,
		},
		{
			Name:  "yidf",
			Money: 123,
		},
		{
			Name:  "sdf",
			Money: 132,
		},
		{
			Name:  "sdfyi",
			Money: 12,
		},
	}
	for i, i2 := range persions {
		fmt.Println(i, i2)
	}
	var total int
	for _, i2 := range persions {
		fmt.Println(i2)
		total += i2.Money
	}

	fmt.Println("total:", total)

	persions[2].Money = 11111111111
	persions[3].Money = 13232323
	for _, i2 := range persions {
		fmt.Println(i2)
		total += i2.Money
	}

}
