// Package  provides ...
package main

import (
  "fmt"
)

func main() {
  // name := "Yusril"; // true
  name := "Arzaqi"; // false

  if name == "Yusril" /* true */ {
    fmt.Println("Hello Yusril"); // print jika name "Yusril"
  } else if name == "Arzaqi" {
    fmt.Println("Hello Arzaqi") // print jika name "Arzaqi"
  } else {
    fmt.Println("Hai, ... ?"); // print jika name tidak "Yusril" atau "Arzaqi"
  }

  if length := len(name); length > 5 {
    fmt.Println("Nama terlalu panjang")
  } else {
     fmt.Println("Nama sudah benar")
  }
}
