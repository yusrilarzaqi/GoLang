package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
	fmt.Println("Terjadi Error", message)
	}
	fmt.Println("Selesai memanggil Function")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
  fmt.Println("Testing")
}
