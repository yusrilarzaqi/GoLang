package main

import "fmt"

func main() {
	var name1 = "Yusril"
	var name2 = "Arzaqi"

	var result = name1 == name2 // false
	fmt.Println(result)
	result = name1 != name2 // true
	fmt.Println(result)

	var num1 = 100
	var num2 = 200
	fmt.Println(num1 > num2)  // false
	fmt.Println(num1 < num2)  // true
	fmt.Println(num1 >= num2) // false
	fmt.Println(num1 <= num2) // true
	fmt.Println(num1 == num2) // false
	fmt.Println(num1 != num2) // true

}
