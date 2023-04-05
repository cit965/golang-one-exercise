package main

import "fmt"

type Person struct {
	Name string
}

type Bus struct {
	P     Person
	Price int
}

func (b *Bus) increaePrice(money int) {
	b.Price += money
}

func main() {
	person := Person{Name: "nange"}
	bus := Bus{
		P:     person,
		Price: 500000,
	}
	fmt.Println(bus)
	fmt.Println(bus.P.Name)
	bus.increaePrice(2)
	fmt.Println(bus.Price)
}
