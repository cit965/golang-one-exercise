package main

import "fmt"

func main() {
	math := 23
	english := 25
	if math > english {
		fmt.Println("nan ge shu xue hen hao")
	} else {
		fmt.Println("nan ge yingyu bu cuo")
	}

	//
	if english == 25 {
		fmt.Println("nan ge yingyu == 25")
	}

	if math < english {
		//
		fmt.Println("<<<<<")
	} else if math == english {
		//
	} else {
		//
		fmt.Println("else")
	}
}
