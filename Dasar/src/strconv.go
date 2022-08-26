package main

import (
	"fmt"
	"strconv"
)

func main() {
  boolean, err := strconv.ParseBool("true")

  if err == nil {
    fmt.Println(boolean)
  } else {
    fmt.Println("Error", err.Error())
  }

  number, err := strconv.ParseInt("100000", 10, 32)
  
  if err == nil {
    fmt.Println(number)
  } else {
    fmt.Println("Error", err.Error())
  }

  value := strconv.FormatInt(10000000, 16)
  fmt.Println(value)

  valueInteger, err := strconv.Atoi("200000000")
  
  if err == nil {
    fmt.Println(valueInteger)
  } else {
    fmt.Println("Error", err.Error())
  }
} 
