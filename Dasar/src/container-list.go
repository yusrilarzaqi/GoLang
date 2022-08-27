package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Yusril")
	data.PushBack("Arzaqi")
	data.PushBack("Bimo")
	data.PushBack("Alamsyah")

	// dari depan ke belakang
	// for element := data.Front(); element != nil; element = element.Next() {
	//   fmt.Println(element.Value)
	// }

	// dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

	// fmt.Println(data.Front().Prev())
	// fmt.Println(data.Back().Next())
	fmt.Println(data.Len())
}
