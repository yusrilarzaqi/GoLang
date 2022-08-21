package main

import "fmt"

func Ups(i int) any {
	switch i {
	case 1:
		return 1
	case 2:
		return true
	default:
		return "Ups"
	}
}

func main() {
	kosong := Ups(4)
	fmt.Println(kosong)

  var data  = Ups(2)
  fmt.Println(data)
}
