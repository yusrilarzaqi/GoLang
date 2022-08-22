package main

import "fmt"

type Address struct {
  City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address) {
  address.Country = "Indonesia"
}

func main() {
  address := Address{
    "Semarang",
    "JawaTengah",
    "",
  }

  fmt.Println(address)

  // sama saja
  // var alamatPointer *Address = &address
  ChangeAddressToIndonesia(&address) // lebih ringkas
  fmt.Println(address) // akan berubah
}
