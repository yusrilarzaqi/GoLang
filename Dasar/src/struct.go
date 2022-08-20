package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	IsMarried       bool
}

func main() {
	var yusril Customer
	yusril.Name = "Yusril Arzaqi"
	yusril.Address = "Indonesia"
	yusril.Age = 18

	fmt.Println(yusril)
	fmt.Println(yusril.Name)
	fmt.Println(yusril.Address)
	fmt.Println(yusril.Age)

	bimo := Customer{
		Name:    "Bimo Alamsyah",
		Address: "Indonesia",
		Age:     17,
	}

	fmt.Println(bimo)
	fmt.Println(bimo.Name)
	fmt.Println(bimo.Address)
	fmt.Println(bimo.Age)

	adam := Customer{"Adam Saputra", "Indonesia", 18, false}

	fmt.Println(adam)
	fmt.Println(adam.Name)
	fmt.Println(adam.Address)
	fmt.Println(adam.Age)
	fmt.Println(adam.IsMarried)
}
