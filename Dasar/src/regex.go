package main

import (
	"fmt"
	"regexp"
)


func main() {
	regex := regexp.MustCompile(`e([a-z])`) 

	fmt.Println(regex.MatchString("eko")) // true
	fmt.Println(regex.MatchString("edo")) // true
	fmt.Println(regex.MatchString("eKo")) // false

  search := regex.FindAllString("eko edo egi ego e1o eto", 10)

  for _, v := range search {
    fmt.Println(v)
  }
}
