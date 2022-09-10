package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	go func() {
		channel <- "Yusril"
		channel <- "Arzaqi"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)

		fmt.Println(cap(channel)) // melihat panjang buffer
		fmt.Println(len(channel)) // melihat jumlah data dibuffer
	}()

	fmt.Println("Selesai")
  time.Sleep(2 * time.Second)
}
