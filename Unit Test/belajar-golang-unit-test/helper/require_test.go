package helper

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
  result := HelloWorld("Yusril")
  require.Equal(t, "Hell Yusril", result)
  fmt.Println("tidak dieksekusi")
}
