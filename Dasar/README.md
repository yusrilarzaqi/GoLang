# Dasar GoLang

## Pengenalan

### Sejarah Go-Lang

- Dibuat di Google menggunakan bahasa pemrograman C.
- Di Rilis ke public sebagai open source pada tahun 2009.
- Go-Lang populer sejak digunakan untuk membuat Docker pada tahun 2011.
- Saat ini mulai banyak teknologi baru yang dibuat menggunakan bahasa Go-Lang dibandingkan bahasa C, seperti Kubernetes, Prometheus, CockroarchDB, dan lain-lain.
- Saat ini mulai populer untuk pembuatan Backend API di Microservices.

### Kenapa Belajar Go-Lang

- Bahasa Go-Lang sangat sederhana, tidak butuh waktu lama untuk mempelajarinya.
- Go-Lang mendukung baik concurrency programming, dimana saat ini kita hidup di zaman multicore processor.
- Go-Lang mendukung garbage collector, sehingga tidak butuh melakukan management memory secara manual seperti di bahasa C.
- Salah satu bahasa pemrograman yang sedang naik daun.

### Software Development Kit

- [Golang](https://golang.org/)
- Download Compiler Go-Lang.
- Install Compiler Go-Lang.
- Cek menggunakan perintah :

```sh
go version
```

### Proses Development Program Go-Lang

![Proses Development Program Go-Lang](./img/Proses-Development-Program-Golang.png)

## Text Editor atau IDE

- Visual Studio Code.
- JetBrains GoLand.

## Program Hello World

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello World")
}
```

## Tipe Data Number

- Ada dua jenis tipe data Number, yaitu :
  - Integer.
  - Floating Point.

| Tipe Data | Nilai Minimum        | Nilai Maksimum      |
| --------- | -------------------- | ------------------- |
| int8      | -128                 | 127                 |
| int16     | -32768               | 32767               |
| int32     | -21474883648         | 21474883647         |
| int64     | -9223372036843775808 | 9223372036843775807 |

###
