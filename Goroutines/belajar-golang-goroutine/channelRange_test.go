package belajargolanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRangeChannel(t *testing.T) {
  channel := make(chan string)  

  go func() {
    for i := 0; i < 10; i++ {
      channel <- "Perulangn " + strconv.Itoa(i)
    }
    close(channel)
  }()

  for data := range channel {
    fmt.Printf("mengirim Data: %v\n", data)
  }
  
  fmt.Println("DONE")
}
