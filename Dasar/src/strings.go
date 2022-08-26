package main

import (
	"fmt"
	"strings"
)

func main() {
  fmt.Println(strings.Contains("Yusril Arzaqi Mantap","Yusril"))
  fmt.Println(strings.Split("Yusril Arzaqi Mantap", " "))
  fmt.Println(strings.ToLower("Yusril Arzaqi"))
  fmt.Println(strings.ToUpper("Yusril Arzaqi"))
  fmt.Println(strings.Trim("                  Yusril Arzaqi           ", " "))
  fmt.Print(strings.ReplaceAll("Yusril Yusril Yusril Yusril Yusril", "Yusril", "Arzaqi"))
}
