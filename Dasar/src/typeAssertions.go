package main

import "fmt"

func random() any {
	return 10
}

func main() {
	// result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	/* resultInt := result.(int) // panic
	fmt.Println(resultInt) */

	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
  case int:
    fmt.Println("Int", value)
  default:
    fmt.Println("Unknown Type : ", value)
	}
}
