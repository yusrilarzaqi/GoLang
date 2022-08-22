package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		"Semarang",
		"Jawa Tengah",
		"Indonesia",
	}

  address2 := &address1
	address2.City = "Demak"
  address3 := &address1 

	fmt.Println(address2)
  
  *address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1) // address1 akan berubah
	fmt.Println(address2) // berubah menyerupai address1
	fmt.Println(address3) // menjadi address1


  // membuat pointer menggunakan &
  /* var address4 *Address = &Address{
		"Semarang",
		"Jawa Tengah",
		"Indonesia",
  } */

  

  // fmt.Println(address4)
}
