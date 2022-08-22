package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
  alamat1 := new(Address)
  alamat2 := alamat1

  alamat2.Country = "Indonesia"
  alamat2.City = "Semarang"

  fmt.Println(alamat1) // alamat 1 akan berubah
  fmt.Println(alamat2)
}
