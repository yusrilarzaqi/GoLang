package belajargolanggoroutine

import (
	"fmt"
	"testing"
)

func TestDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	counter := 0
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

  sec := 0;
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 1", data)
			counter++
		default:
			fmt.Println("Menunggu Data", sec)
      sec++
		}

		if counter == 2 {
			break
		}
	}
  fmt.Printf("sec: %v\n", sec)
}
