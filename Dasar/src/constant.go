package main

import (
	"fmt"
)

func main() {
	// const firstName string = "Yusril"
	// const lastName = "Arzaqi"
	// const value = 1000

	const (
		firstName = "Yusril"
		lastName  = "Arzaqi"
		value     = 1000
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	/* ERROR
	 * firstName = "tidak bisa diubah"
	 * lastName = "tidak bisa diubah"
	 */
}
