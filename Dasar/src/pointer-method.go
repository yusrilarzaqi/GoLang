package main

import "fmt"

type Man struct {
  Name string
}

func (man *Man) Married() {
  man.Name = "Mr, " + man.Name
}

func main() {
  yusril := Man{"Yusril"}
  fmt.Println(yusril)
  yusril.Married()
  fmt.Println(yusril)
}
