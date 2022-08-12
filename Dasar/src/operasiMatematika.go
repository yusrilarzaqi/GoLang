package main

import "fmt"

func main() {
  var a = 10
  var b = 10
  var result = a + b

  fmt.Println(result)
  
  var i = 10;
  i += 10 // i = i + 10
  fmt.Println(i) // 20

  i++ // i = i + 1
  fmt.Println(i) // 21

  var positive = +1000
  var negative = -1000
  fmt.Println(positive)
  fmt.Println(negative)
}
