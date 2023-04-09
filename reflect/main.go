package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{Name: "Tom", Age: 18}

	// 获取变量的类型信息
	fmt.Println(reflect.TypeOf(user)) // 输出：main.User

	// 获取变量的值信息
	fmt.Println(reflect.ValueOf(user)) // 输出：{Tom 18}

	// 获取变量的名称和类型信息
	t := reflect.TypeOf(user)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%s %s\n", field.Name, field.Type)
	}
	// 输出：
	// Name string
	// Age int

	// 修改变量的值
	v := reflect.ValueOf(&user).Elem()
	v.FieldByName("Name").SetString("Jerry")
	v.FieldByName("Age").SetInt(20)
	fmt.Println(user) // 输出：{Jerry 20}
}

//以上代码中，我们定义了一个 User 结构体，然后使用反射获取了它的类型信息和值信息，以及其中的字段名称和类型信息。最后，我们还演示了如何使用反射来修改变量的值。
