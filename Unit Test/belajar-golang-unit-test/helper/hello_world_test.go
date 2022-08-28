package helper

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
  result := HelloWorld("Yusril")

  if result != "Hello Yusril" {
    // unit test failed
    panic("Result is not Hello Yusril")
  }
}

func TestHelloWorldArzaqi(t *testing.T) {
  result := HelloWorld("Arzaqi")

  if result != "Hello Arzaqi" {
    // unit test failed
    panic("Result is not Hello Arzaqi")
  }
}
