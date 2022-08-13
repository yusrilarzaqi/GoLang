package main
import "fmt"

func main() {
  var person /* map[string]string */ = map[string]string {
    "firstName" : "Yusril",
    "lastName" : "Arzaqi",
    "address": "Semarang",
  };

  person["title"] = "Mahasiswa"

  fmt.Println(person) // map[address:Semarang firstName:Yusril lastName:Arzaqi title:Mahasiswa]
  fmt.Println(person["firstName"]) // Yusril
  fmt.Println(person["lastName"]) // Arzaqi
  fmt.Println(person["address"]) // Semarang
  fmt.Println(person["title"]) // Mahasiswa

  book := make(map[string]string) // membuat map kosong
  book["title"] = "Buku Go-Lang" 
  book["author"] = "Yusril Arzaqi"
  book["wrong"] = "Ups"

  // sebelum dihapus
  fmt.Println("Sebelum Dihapus")
  fmt.Println(len(book))
  fmt.Println(book)
  
  delete(book, "wrong") // menghapus key "wrong"

  // setelah dihapus
  fmt.Println("Setelah Dihapus")
  fmt.Println(book)
  fmt.Println(len(book))
}
