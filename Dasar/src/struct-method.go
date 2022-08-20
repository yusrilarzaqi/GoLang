package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	IsMarried     bool
}

// Memakai function biasa
// func sayHi(customer Customer, name string) {
//   fmt.Println("Hello,", name, "My name is", customer.Name, "salam kenal.")
// }

// menggunakan struct method
func (customer Customer) sayHi(name string) {
  fmt.Println("Hello,", name, "My Name is", customer.Name, "salam kenal.")
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, My Neme is", customer.Name)
}


func main() {
  yusril := Customer{Name: "Yusril Arzaqi"}
  yusril.sayHello()
  yusril.sayHi("Bimo")
}
