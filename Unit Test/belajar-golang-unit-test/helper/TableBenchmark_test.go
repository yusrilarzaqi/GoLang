package helper

import (
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name, request string }{
		{
			name:    "HelloWorld(Yusril)",
			request: "Yusril",
		},
		{
			name:    "HelloWorld(Arzaqi)",
			request: "Arzaqi",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
