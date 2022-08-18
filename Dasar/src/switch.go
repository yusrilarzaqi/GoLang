package main

import "fmt"

func main() {
	name := "bimo"

	switch name {
	case "Yusril":
		fmt.Println("Hello Yusril")
	case "Adam":
		fmt.Println("Hello Adam")
	default:
		fmt.Println("Hi, Boleh Kenalan ?")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
  case length > 5:
    fmt.Println("Nama Lumayan Panjang")
  default :
    fmt.Println("Nama Sudah Benar")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama Terlalu Panjang")
	// case false:
	// 	fmt.Println("Nama Sudah Benar")
	// }
}
