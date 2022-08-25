package helper

var version = "1.0.0" // tidak bisa diakses dari luar
var Application = "golang" 

// Tidak bisa diakses dari luar package
func sayGoodBye(name string) string {
  return "Hello" + name
}

func SayHello(name string) string {
  return "Hello " + name
}
