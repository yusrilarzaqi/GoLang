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

| Tipe Data |    Nilai Minimum     |    Nilai Maksimum    |
| :-------: | :------------------: | :------------------: |
|   int8    |         -128         |         127          |
|   int16   |        -32768        |        32767         |
|   int32   |     -21474883648     |     21474883647      |
|   int64   | -9223372036843775808 | 9223372036843775807  |
|   uint8   |          0           |         255          |
|  uint16   |          0           |        65535         |
|  uint32   |          0           |      4294967295      |
|  uint64   |          0           | 18446744073709551615 |

### Tipe Data Floating Point

| Tipe Data  |                     Nilai Minimum                      | Nilai Maksimum |
| :--------: | :----------------------------------------------------: | :------------: |
|  float32   |                        -3.4e+38                        |    3.4e+38     |
|  float64   |                       -1.7e+308                        |   +1.7e+308    |
| complex64  | complex numbers with float32 real and imaginary parts. |       -        |
| complex128 | complex numbers with float64 real and imaginary parts. |       -        |

### Alias

| Tipe Data |     Alias      |
| :-------: | :------------: |
|   byte    |     uint8      |
|   rune    |     int32      |
|    int    | Minimal int32  |
|   uint    | Minimal uint32 |

### Kode Program Number

```go
package main

import "fmt"

func main() {
  fmt.Println("Satu = ", 1)
  fmt.Println("Dua = ", 2)
  fmt.Println("Tiga Koma Lima = ", 3.5)
}
```

## Tipe Data Boolean

- Tipe data boolean adalah tipe data yang memiliki dua nilai, yaitu benar atau salah.
- Di Go-Lang, tipe data boolean direpresentasikan menggunakan kata kunci `bool`.

### Kode : Boolean

```go
package main

import "fmt"

func main() {
  fmt.Println("Benar = ", true)
  fmt.Println("Salah = ", false)
}
```

##
