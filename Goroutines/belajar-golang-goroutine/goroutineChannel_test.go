package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	/* channel <- "Yusril"
	result := <- channel
	fmt.Printf("result: %v\n", result)
	fmt.Println(<- channel) */

	go func() {
		time.Sleep(1 * time.Second)
		channel <- "Yusril Arzaqi"
		fmt.Println("Selesai Mengirim Data ke Channel.")
	}()

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Yusril Arzaqi"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	close(channel)
}
