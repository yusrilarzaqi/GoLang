package main

import "fmt"

func main() {
  name := "Yusril"
  counter := 0
  
  increment := func() {
    const name = "Dimas"
    fmt.Println("Increment")
    counter++; // counter akan bisa di akses
  }

  increment()
  increment()
  fmt.Println(counter)
  fmt.Println(name)
}
