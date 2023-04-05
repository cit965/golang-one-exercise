package main

import (
	"fmt"
	"github.com/mouuui/init/ano"
)

var xx string

// init() is ran before the main() function
func init() {
	xx = "nihao"
}

func main() {
	fmt.Println(xx)
	fmt.Println(ano.F)
}
