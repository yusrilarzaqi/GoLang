package helper

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Yusril")
	assert.Equal(t, "Hello Yusril", result, "Result must be Hello Yusril")
	fmt.Println("Dieksekusi")
}
