package helper

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
  t.Run("Yusril", func(t *testing.T) {
    result := HelloWorld("Yusril")
    require.Equal(t, "Hello Yusril", result)
  })

  t.Run("Arzaqi", func(t *testing.T) {
    result := HelloWorld("Arzaqi")
    require.Equal(t, "Hello Arzaqi", result)
  })
}
