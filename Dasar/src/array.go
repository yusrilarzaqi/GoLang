package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Yusril"
	names[1] = "Arzaqi"

	fmt.Println(names[0]) // Yusril
	fmt.Println(names[1]) // Arzaqi
	fmt.Println(names)    // [Yusril Arzaqi]

	var values = [3]int{
		95,
		90,
		80,
	}
	fmt.Println(values) //  [95, 90, 80]

	fmt.Println(len(names)) // 2
	fmt.Println(len(values)) // 3

  var test [10]int

  fmt.Println(len(test)) // 10

}
