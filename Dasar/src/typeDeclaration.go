package main

import "fmt"

func main() {
  type NoKTP string
  type Married bool

  var ktpYus NoKTP = "101001931013"
  var marriedStatus Married = true
  fmt.Println(ktpYus)
  fmt.Println(marriedStatus)
  fmt.Println(NoKTP("22121313131"))
}
