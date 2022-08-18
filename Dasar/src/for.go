package main

import (
	"fmt"
)

func main() {
	// counter := 1
	//
	// for counter <= 10 {
	//   fmt.Println("Perulangan Ke : ", counter)
	//   counter++;
	// }

	// for counter := 1; counter <= 10; counter++ {
	//    fmt.Println("Perulangan Ke : ", counter)
	// }

	names := []string{
		"Yusril",
		"Arzaqi",
		"Bimo",
		"Adam",
	}

	for _, name := range names {
		// fmt.Println("Index", index, "=", name)
		fmt.Println(name)
	}
	// for i := 0; i < len(names); i++ {
	//   fmt.Println(names[i])
	// }

	person := map[string]string{
		"name":  "Yusril",
		"kelas": "TKJ2",
		"nim":   "185512",
	}

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}
