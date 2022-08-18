package main

import "fmt"

func factoricalLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * (factorialRecursive(value - 1))
	}
}

func main() {
	fmt.Println("faktorical menggunakan loop", factoricalLoop(5))
	fmt.Println("faktorical menggunakan Recursive", factorialRecursive(5))
	fmt.Println("faktorical menggunakan manual", 1*2*3*4*5)
}
