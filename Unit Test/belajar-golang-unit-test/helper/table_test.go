package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelloWorldTable(t *testing.T) {
	tests := [...]struct {
		name, request, expected string
	}{
		{
			name:     "HelloWold(Yusril)",
			request:  "Yusril",
			expected: "Hello Yusril",
		},
    {
      name: "HelloWorld(Bimo)",
      request: "Bimo",
      expected: "Hello Bimo",
    },
    {
      name: "HelloWorld(Arzaqi)",
      request: "Arzaqi",
      expected: "Hello Arzaqi",
    },
	}

  for _, test := range tests {
    t.Run(test.name, func(t *testing.T) {
      result := HelloWorld(test.request)
      require.Equal(t, test.expected, result)
    })
  }
}
