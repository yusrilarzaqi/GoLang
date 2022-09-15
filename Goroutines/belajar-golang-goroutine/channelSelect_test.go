package belajargolanggoroutine

import (
	"fmt"
	"testing"
)

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	counter := 0
  defer close(channel1)
  defer close(channel2)

  go GiveMeResponse(channel1)
  go GiveMeResponse(channel2)

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}

  fmt.Printf("counter: %v\n", counter)
}
