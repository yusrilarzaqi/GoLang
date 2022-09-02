package helper

import "testing"

func BenchmarkSub(b *testing.B) {
  b.Run("Yusril", func(b *testing.B) {
    for i := 0; i < b.N ; i++ {
      HelloWorld("Yusril")
    }
  })
  
  b.Run("Arzaqi", func(b *testing.B) {
    for i := 0; i < b.N ; i++ {
      HelloWorld("Arzaqi")
    }
  })
}
