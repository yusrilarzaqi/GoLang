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

### Kode : Tipe Data Number

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

### Kode : Tipe Data Boolean

```go
package main

import "fmt"

func main() {
  fmt.Println("Benar = ", true)
  fmt.Println("Salah = ", false)
}
```

## Tipe Data String

- String ada tipe data kumpulan karakter.
- Jumlah karakter di dalam String bisa nol sampai tidak terhingga.
- Tipe data String di Go-Lang direpresentasikan dengan kata kunci `string`.
- Nilai data String di Go-Lang selalu diawali dengan karakter `"` (petik dua) dan diakhiri dengan karakter `"` (petik dua).

### Kode : Tipe Data String

```go
package main

import "fmt"

func main() {
  fmt.Println("Yusril")
  fmt.Println("Yusril Arzaqi")
  fmt.Println("Yusril Arzaqi memetik bunga")
}
```

### Fucntion untuk String

| Function           | Keterangan                                     |
| ------------------ | ---------------------------------------------- |
| `len("string)`     | Menghitung jumlah karakter di String           |
| `"string"[number]` | Mengambil karakter pada posisi yang ditentukan |

### Kode : Function untuk String

```go
package main

import "fmt"

func main() {
  fmt.Println("Yusril")
  fmt.Println("Yusril Arzaqi")
  fmt.Println("Yusril Arzaqi memetik bunga")
  fmt.Println(len("Yusril Arzaqi memetik bunga"))
  fmt.Println("Yusril Arzaqi memetik bunga"[10])
}
```

## Variable

- Variable adalah tempat untuk menyimpan data.
- Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau.
- Di Go-Lang Variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan yang berbeda-beda jenis, kita harus membuat beberapa variable.
- Untuk membuat variable, kita bisa menggunakan kata kunci `var`, lalu diikuti dengan nama variable dan tipe datanya.

### Kode : Variable

```go
func main() {
  var name string;

  name = "Yusril Arzaqi"
  fmt.Println(name)

  name = "Bimo Alamsyah"
  fmt.Println(name)
}
```

### Tipe Data Variable

- Saat kita membuat variable, maka kita wajib menyabutkan tipe data variable tersebut.
- Namun jika kita langsung menginisialisasikan data pada variable nya, maka tidak wajib menyebutkan tipe data variablenya.

```go
var friendName = "Adam Saptura";
fmt.Println(friendName)

var age int8 = 18;
fmt.Println(age)
```

### Kata Kunci Var

- Di Go-Lang kata kunci `var` saat membuat variable tidak lah wajib.
- Asalkan saat membuat variable kita langsung menginisialisasi datanya.
- Agar tidak perlu menggunakan kata kunci `var`, kita perlu menggunakan kata kunci `:=` saat menginisialisasikan data pada variable tersebut.

### kode : Kata Kunci Var

```go
author := "Dimas Rafif Fathony"
fmt.Println(author)
```

### Deklarasi Multiple Variable

- Di Go-Lang kita bisa membuat variable secara sekaligus banyak.
- Code yang dibuat sakan lebih bagus dan mudah dibaca.

```go
func main() {
  var (
    firstName = "Yusril"
    lastName =  "Arzaqi"
  )

  fmt.Println(firstName)
  fmt.Println(lastName)
}
```

## Constant

- Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai.
- Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah `const`, bukan `var`.
- Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya.

### Kode : Constant

```go
func main() {
  const firstName string = "Yusril"
  const lastName = "Arzaqi"
  const value = 1000

  fmt.Println(firstName)
  fmt.Println(lastName)
  fmt.Println(value)

  /* ERROR
   * firstName = "tidak bisa diubah"
   * lastName = "tidak bisa diubah"
   */
}
```

### Deklarasi Multiple Constant

- Sama seperti Variabel, di Go-Lang juga kita bisa membuat constant secara sekaligus banyak.

```go
const (
  firstName = "Yusril"
  lastName         = "Arzaqi"
  value            = 1000
)
```

## Konversi Tipe Data

- Di Go-Lang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain.
- Misal kita ingin mengkonversi tipe data `int32` ke `int64`, dan lain-lain.

```go
func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
  fmt.Println(nilai8)

  var name = "Yusril arzaqi"
  var e byte = name[0]
  var eString = string(e)

  fmt.Println(name)
  fmt.Println(eString)
}
```

##
