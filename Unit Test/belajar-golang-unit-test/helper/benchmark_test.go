package helper

import "testing"

func BenchmarkYusril(b *testing.B) {
  for i := 0; i < b.N; i++ {
    HelloWorld("Yusril")
  }
}

func BenchmarkArzaqi(b *testing.B) {
  for i := 0; i < b.N; i++ {
    HelloWorld("Arzaqi")
  }
}
