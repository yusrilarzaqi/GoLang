package helper

import (
	"runtime"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
  if runtime.GOOS == "linux" {
    t.Skip("Unit test tidak bisa jalan di Linux");
  }

  result := HelloWorld("Yusril")
  require.Equal(t, "Hello Yusril", result)
}
