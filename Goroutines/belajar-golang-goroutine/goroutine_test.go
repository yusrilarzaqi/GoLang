package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func SayHello(name string) {
  fmt.Println("Hello " + name)
}

func TestCreateGoroutine(t *testing.T) {
  go SayHello("Yusril")
  // SayHello("Yusril")
  fmt.Println("Ups")

  time.Sleep(1 * time.Second)
}
