package main

import "fmt"

func main() {
  var nilaiAkhir = 90;
  var absensi = 80;

  var lulusNilaiAkhir bool = nilaiAkhir >= 80; // true
  var lulusAbsensi bool = absensi >= 80; // true 

  var lulus bool = lulusAbsensi && lulusNilaiAkhir; // true

  fmt.Println(lulus)
  fmt.Println(nilaiAkhir >= 80 && absensi >= 80)
}
