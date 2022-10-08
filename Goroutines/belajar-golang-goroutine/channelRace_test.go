package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceChannel(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				fmt.Printf("x: %v\n", x)
				x++
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("counter : %v\n", x)

}
