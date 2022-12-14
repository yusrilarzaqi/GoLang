# Modules

## Sebelum Belajar Materi ini

- Go-Lang Dasar

## Agenda

- Mengenal Go Modules.
- Membuat Modules.
- Rilis Module.
- Menambah Dependency.
- Upgrade Module.
- Upgrade Dependency.

## Pengenalan Go-Modules

### Sebelum Ada Go Modules

- Saat kita membuat aplikasi, biasanya akan menggunakan library atau dependency dari project lain.
- Sebelum ada Go Modules, management untuk dependency sangat sulit dilakukan di Go-Lang.
- Biasanya kita akan meng-copy semua source code library dengan library orang lain.
- Atau biasanya library orang lain akan kita save di `GOPATH` directory, namun hal ini akan sulit jika ternyata beberapa aplikasi butuh library yang sama dengan versi yang berbeda.

### Go-Modules

- Go-Modules mulai dikenal di Go-Lang 1.11 dan 1.12.
- Dengan Go-Modules, kita bisa membuat project dengan mudah dan dependency management yang sangat mudah.

## Membuat Module

- Untuk membuat module baru, kita bisa menggunakan perintah `go mod init nama-module`.
- Go-Lang akan secara otomatis membuat file `go.mod` yang berisikan informasi nama module dan juga versi Go-Lang yang kita gunakan Go-Lang yang kita gunakan.

### Rilis Module

- Go-Lang terintegrasi baik dengan Git.
- Untuk merilis module, kita hanya perlu membuat Tag di Git.
- [go-say-hello](https://github.com/arzaqiyusril/go-say-hello)

## Menambah Dependency

`go get nama-module`

### Cara menggunakan module dari github

```go
package main

import (
	"fmt"
	go_say_hello "github.com/arzaqiyusril/go-say-hello"
)

func main() {
	fmt.Println(go_say_hello.SayHello("Yusril"))
}
```

## Upgrade Module

- Untuk melakukan upgrade module, kita hanya cukup membuat tag baru di Git.

## Upgrade Dependency

- Untuk upgrade dependency ke versi terbaru, kita bisa mengubah isi `go.mod`, lalu mengubah tagnya menjadi tag terbaru.
- Untuk download versi terbaru, gunakan perintah `go get`.

## Major Upgrade

- Major upgrade biasanya terjadi dikarenakan ada perubahan pada isi kode program kita, sehingga membuatnya tidak backward compatible.
- Sebaliknya hal ini sebisa mungkin dihindari.
- Namun jika tidak bisa dihindari, strategy terbaik adalah merubah nama module.
