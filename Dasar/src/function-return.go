package main

import "fmt"

func sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}
func getHello(name string) string {
	if name == "" {
		return "Hello Anda Siapa"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Yusril")
	fmt.Println(result)
  
	result = getHello("")
	fmt.Println(result)

	fmt.Println(sum([]int{10, 20, 30, 40}))
}
