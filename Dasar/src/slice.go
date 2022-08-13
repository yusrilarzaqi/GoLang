package main

import "fmt"

func main() {
	names := [...]string{
		"Yusril Arzaqi",
		"Bimo Alamsyah",
		"Adam Saptura",
		"Dimas Rafif",
		"Irfan",
	}
	slice := names[1:4]

	fmt.Println(slice[0])   // Bimo Alamsyah
	fmt.Println(slice[1])   // Adam Saptura
	fmt.Println(len(names)) // 5
	fmt.Println(cap(slice))

	days := [...]string{
		"Senin", "Selasa", "Rabu", "Kamis", "Jumad", "Sabtu", "Minggu",
	}

	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1) // [Sabtu Baru Minggu Baru]
	fmt.Println(days)      // [Senin Selasa Rabu Kamis Jumad Sabtu Baru Minggu Baru]

	daySlice2 := append(daySlice1, "Libur Baru") // membuat dan menambah slice baru dari daySlice1
	daySlice2[0] = "Ups"                         // mengganti Sabtu Baru menjadi "Ubs"
	fmt.Println(daySlice2)                       // [Ups Minggu Baru Libur Baru]
	fmt.Println(days)                            // [Senin Selasa Rabu Kamis Jumad Sabtu Baru Minggu Baru]

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Yusril"
	newSlice[1] = "Arzaqi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(toSlice)

  iniArray := [5]int{1, 2, 3, 4, 5}
  iniSlice := []int{1,2,3 ,4, 5 }

  fmt.Println(iniArray)
  fmt.Println(iniSlice)
}
