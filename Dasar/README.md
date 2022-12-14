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

## Type Declaration

- Type Declaration adalah kemampuan membuat ulang tipe data dari tipe data yang sudah ada.
- Type Declaration biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti.

```go
func main() {
  type NoKTP string
  type Married bool

  var ktpYus NoKTP = "101001931013"
  var marriedStatus Married = true
  fmt.Println(ktpYus)
  fmt.Println(marriedStatus)
  fmt.Println(NoKTP("22121313131"))
}
```

## Operasi Matematika

| Operator |   Keterangan    |
| :------: | :-------------: |
|    +     |   Penjumlahan   |
|    -     |   Pengurangan   |
|    \*    |    Perkalian    |
|    /     |   Pembagaian    |
|    %     | Sisa Pembagaian |

### Kode : Operasi Matematika

```go
func main() {
  var a = 10
  var b = 10
  var result = a + b

  fmt.Println(result) // 20
}
```

### Augmented Assignments

| Operasi Matematika | Augmented Assignments |
| :----------------: | :-------------------: |
|    `a = a + 10`    |       `a += 10`       |
|    `a = a - 10`    |       `a -= 10`       |
|    `a = a * 10`    |       `a *= 10`       |
|    `a = a / 10`    |       `a /= 10`       |
|    `a = a % 10`    |       `a %= 10`       |

```go
func main() {
  var i = 10;
  i += 10
  fmt.Println(i)
}
```

### Unary Operator

| Operator |    Keterangan     |
| :------: | :---------------: |
|   `++`   |    `a = a + 1`    |
|   `--`   |    `a = a - 1`    |
|   `-`    |     Negative      |
|   `+`    |     Positive      |
|   `!`    | Boolean Kebalikan |

### Kode : Unary Operator

```go
i++ // i = i + 1
fmt.Println(i) // 21

var positive = +1000
var negative = -1000
fmt.Println(positive)
fmt.Println(negative)
```

## Operasi Perbandingan

- Operasi perbandingan adalah operasi untuk membandingkan dua buah data.
- Operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah).
- Jika hasil operasinya adalah benar, maka nilainya adalah `true`.
- Jika hasil operasinya adalah salah, maka nilainya adalah `false`.

| Operator |       Keterangan        |
| :------: | :---------------------: |
|   `>`    |       Lebih Dari        |
|   `<`    |       Kurang Dari       |
|   `>=`   | Lebih Dari Sama Dengan  |
|   `<=`   | Kurang Dari Sama Dengan |
|   `==`   |       Sama Dengan       |
|   `!=`   |    Tidak Sama Dengan    |

## Kode : Operasi Perbandingan

```go
func main() {
  var name1 = "Yusril";
  var name2 = "Arzaqi";

  var result bool = name1 == name2; // false
  fmt.Println(result)
  result = name1 != name2; // true
  fmt.Println(result)

  var num1 = 100
  var num2 = 200
  fmt.Println(num1 > num2) // false
  fmt.Println(num1 < num2) // true
  fmt.Println(num1 >= num2) // false
  fmt.Println(num1 <= num2) // true
  fmt.Println(num1 == num2) // false
  fmt.Println(num1 != num2) // true
```

## Operasi Boolean

| Operator | Keterangan |
| :------: | :--------: |
|   `&&`   |    Dan     |
|  `\|\|`  |    Atau    |
|   `!`    | Kebalikan  |

### Operasi &&

| Nilai 1 | Operator | Nilai 2 | Hasil   |
| :-----: | :------: | :-----: | ------- |
| `true`  |   `&&`   | `true`  | `true`  |
| `true`  |   `&&`   | `false` | `false` |
| `false` |   `&&`   | `true`  | `false` |
| `false` |   `&&`   | `false` | `false` |

### Operator ||

| Nilai 1 | Operator | Nilai 2 | Hasil   |
| :-----: | :------: | :-----: | ------- |
| `true`  |  `\|\|`  | `true`  | `true`  |
| `true`  |  `\|\|`  | `false` | `true`  |
| `false` |  `\|\|`  | `true`  | `true`  |
| `false` |  `\|\|`  | `false` | `false` |

### Operator ! (not)

| Operator | Nilai 2 |  Hasil  |
| :------: | :-----: | :-----: |
|   `!`    | `true`  | `fasle` |
|   `!`    | `false` | `true`  |

### Kode : Operasi Boolean

```go
func main() {
  var nilaiAkhir = 90;
  var absensi = 80;

  var lulusAbsensi bool = absensi > 80; // false
  var lulusNilaiAkhir bool = nilaiAkhir >= 80; // true
  var lulusAbsensi bool = absensi >= 80; // true

  var lulus bool = lulusAbsensi && lulusNilaiAkhir; // true

  fmt.Println(lulus)
  fmt.Println(nilaiAkhir >= 80 && absensi >= 80) // true
}
```

## Tipe Data Array

- Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama.
- Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh Array tersebut.
- Daya tampung Array tidak bisa bertambah setelah Array dibuat.

### Index di Array

| Index |  Data  |
| :---: | :----: |
|   0   | Yusril |
|   1   | Arzaqi |

### Kode : Tipe Data Array

```go
func main() {
	var names [2]string
	names[0] = "Yusril"
	names[1] = "Arzaqi"

	fmt.Println(names[0]) // yusril
	fmt.Println(names[1]) // arzaqi
}
```

### Membuat Array Langsung

- Di Go-Lang kita juga bisa membuat Array secara langsung saat deklarasi variable.

```go
var values = [3]int{
  95,
  90,
  80,
};
fmt.Println(values) //  [95, 90, 80]
```

### Function Array

|        Operasi         |           Keterangan            |
| :--------------------: | :-----------------------------: |
|      `len(array)`      | Untuk mendapatkan panjang Array |
|     `array[index]`     |  Mendapat data di posisi index  |
| `array[index] = value` |  Mengubah data di posisi index  |

```go
fmt.Println(len(names))
fmt.Println(len(values))

var test [9]int
fmt.Println(len(test)) // 10
```

## Tipe Data Slice

- Tipe data Slice adalah potongan dari data Array.
- Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah.
- Slice dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagaian atau seluruh data di Array.
- Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity.
- Pointer adalah penunjuk data pertama di array pada slice.
- Length adalah panjang dari slice, dan
- Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity.

### Detail Tipe Data Slice

- Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity.
- Pointer adalah penunjuk data pertama di array para slice.
- Length adalah penunjuk data pertama di array para slice.
- Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity.

### Membuat Slice Dari Array

|   Membuat Slice   |                               Keterangan                               |
| :---------------: | :--------------------------------------------------------------------: |
| `array[low:high]` |  Membuat slice dari array dimulai index low sampai index sebelum high  |
|   `array[low:]`   | Membuat slice dari array dimulai index low sampai index akhir di array |
|  `array[:high]`   |   Membuat slice dari array dimulai index 0 sampai index sebelum high   |
|    `array[:]`     | Membuat slice dari array dimulai index 0 sampai index akhir dari array |

### Slice dan Array

![Slice dan Array](./img/SlicedanArray.png)

### Kode : Tipe Data Slice

```go
func main() {
  names := [...]string{
    "Yusril Arzaqi",
    "Bimo Alamsyah",
    "Adam Saptura",
    "Dimas Rafif",
    "Irfan",
  }
  slice := names[1:4]

  fmt.Println(slice[0]) // Bimo Alamsyah
  fmt.Println(slice[1]) // Adam Saptura
  fmt.Println(len(names)) // 5
}
```

### Function Slice

|               Operasi                |                                    Keterangan                                    |
| :----------------------------------: | :------------------------------------------------------------------------------: |
|             `len(slice)`             |                         Untuk mendapatkan panjang slice                          |
|             `cap(slice)`             |                           Untuk mendapatkan kapasitas                            |
|        `append(slice, data)`         | Memubat slice baru dengan nemambah data ke posisi terakhir slice, jika kapasitas |
| `make([]TypeData, length, capacity)` |                                Membuat slice baru                                |
|     `copy(destination, source)`      |                    Menyalin slice dari source ke destination                     |

### Kode : Append Slice

```go
  daySlice2 := append(daySlice1, "Libur Baru") // membuat dan menambah slice baru dari daySlice1
  daySlice2[0] = "Ups" // mengganti Sabtu Baru menjadi "Ubs"
  fmt.Println(daySlice2) // [Ups Minggu Baru Libur Baru]
  fmt.Println(days) // [Senin Selasa Rabu Kamis Jumad Sabtu Baru Minggu Baru]

```

### Kode : Make Slice

```go
newSlice := make([]string, 2, 5)
newSlice[0] = "Yusril"
newSlice[1] = "Arzaqi"

fmt.Println(newSlice)
fmt.Println(len(newSlice))
fmt.Println(cap(newSlice))
```

### Kode : Copy Slice

```go
fromSlice := days[:]
toSlice := make([]string, len(fromSlice), cap(fromSlice))
copy(toSlice, fromSlice)

fmt.Println(toSlice)
```

### Hati - Hati Saat Membuat Array

- Saat membuat Array, kita harus berhati-hati, jika salah, maka yang kita buat bukanlah array, melainkan slice.

```go
iniArray := [5]int{1, 2, 3, 4, 5}
iniSlice := []int{1,2,3 ,4, 5 }

fmt.Println(iniArray)
fmt.Println(iniSlice)
```

## Tipe Data Map

- Pada Array atau Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0.
- Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan.
- Sederhananya, Map adalah tipe data dan kumpulan key-value (kata kunci - nilai), dimana kata kunci bersifat unik, tidak boleh sama.
- Berbeda dengan Array dan Slice, jumlah data yang kita masukan ke dalam Map boleh sebanyak-banyaknya, asalkan kata kuncinya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru.

### Kode : Tipe Data Map

```go
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
}
```

### Function Map

|            Operasi            |              Keterangan              |
| :---------------------------: | :----------------------------------: |
|          `len(map)`           | Untuk mendapatkan jumlah data di map |
|          `map[key]`           |   Mengambil data di map dengan key   |
|      `map[key] = value`       |   Mengubah data di map dengan key    |
| `make(map[TypeKey]TypeValue)` |           Membuat map baru           |
|      `delete(map, key)`       |   Menghapus data di map dengan key   |

### Kode : Function Map

```go
func main() {
  book := make(map[string]string) // membuat map kosong
  book["title"] = "Buku Go-Lang"
  book["author"] = "Yusril Arzaqi"
  book["wrong"] = "Ups"

  delete(book, "wrong") // menghapus key "wrong"

  fmt.Println(book)
}
```

## If Expression

- If adalah salah satu kata kunci yang digunakan untuk percabangan.
- Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi.
- Hampir di semua bahasa pemrograman mendukung if expression.

### Kode : if Expression

```go
func main() {
    name := "Yusril"

    if name == "Yusril" {
        fmt.Println("Hello Yusril")
    }
}
```

### Else Expression

- Blok `if` akan dieksekusi ketika kondisi `if` bernilai `true`.
- Kadang kita ingin melakukan eksekusi program tertentu jika kondisi `if` bernilai `false`.
- Hal ini bisa dilakukan menggunakan `else` expression.

### Kode : Else Expression

```go
func main() {
  name := "Yusril";

  if name == "Yusril" /* true */ {
    fmt.Println("Hello Yusril");
  } else {
    fmt.Println("Hai, ... ?")
  }
}
```

### Else If Expression

- Kadang dalam if, kita butuh membuat beberapa kondisi.
- Kasus seperti ini, kita bisa menggunakan `else if` expression.

### Kode : Else If Expression

```go
func main() {
  // name := "Yusril"; // true
  name := "Arzaqi"; // false

  if name == "Yusril" /* true */ {
    fmt.Println("Hello Yusril"); // print jika name "Yusril"
  } else if name == "Arzaqi" {
    fmt.Println("Hello Arzaqi") // print jika name "Arzaqi"
  } else {
    fmt.Println("Hai, ... ?"); // print jika name tidak "Yusril" atau "Arzaqi"
  }
}
```

### If dengan Shot Statement

- `if` mendukung shor statement sebelum kondisi.
- Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi.

### Kode : if short statement

```go
func main() {
    if length := len(name); length > 5 {
        fmt.Println("Nama Terlalu panjang");
    } else {
        fmt.Println("Nama sudah benar")
    }
}
```

## Switch Expression

- Selain `if` expression, untuk melakukan percabangan, kita juga bisa menggunakan Switch Expression.
- Switch expression sangat sederhana dibandingkan `if`.
- Bisasanya `switch` expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable.

### Kode : Switch

```go
func main() {
	name := "Yusril"

	switch name {
	case "Yusril":
		fmt.Println("Hello Yusril")
	case "Adam":
		fmt.Println("Hello Adam")
	default:
		fmt.Println("Hi, Boleh Kenalan ?")
	}
}
```

### Switch dengan Short Statement

- Sama dengan `if`, `switch` juga mendukung short statement sebelum variable yang akan di cek kondisinya.

### Kode : Switch Short Hand

```go
func main() {
  name := "Yusril"

  switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}
}
```

### Switch Tanpa Kondisi

- Kondisi di `switch` expression tidak wajib.
- Jika kita tidak menggunakan, kondisi di `switch` expression, kita bisa menambahkan kondisi tersebut di setiap case nya.

### Kode Switch Tanpa Kondisi

```go
func main() {
  name := "Yusril"
  length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
  case length > 5:
    fmt.Println("Nama Lumayan Panjang")
  default :
    fmt.Println("Nama Sudah Panjang")
	}

}
```

## For Loops

- Falam Bahasa Pemrograman, bisasanya ada fitur yang bernama perulangan.
- Salah satu fitur perulangan adalah for loops.

### Kode For Loops

```go
func main() {
  counter := 1

  for counter <= 10 {
    fmt.Println("Perulangan Ke : ", counter)
    counter++;
  }
}
```

### For dengan Statement

- Dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa tambahkan di for.
- Init statement, yaitu sebelum for dieksekusi.
- Post statement, yaitu statement yang akan selalu dieksekusi di akhir tiap perulangan.

### Kode : For use Statement

```go
func main() {
	for counter := 1; counter <= 10; counter++ {
    fmt.Println("Perulangan Ke : ", counter)
	}
}
```

### For Range

- For bisa digunakan untuk melakukan iterasi terhadap semua data collection.
- Data collection contohnya Array, Slice dan Map

### Kode : For Range

**menggunakan for loop biasa**

```go
func main() {
  names := []string {
    "Yusril",
    "Arzaqi",
    "Bimo",
    "Adam",
  }

  for i := 0; i < len(names); i++ {
    fmt.Println(names[i])
  }
}
```

---

**Menggunakan Range**

```go
  func main() {
  names := []string {
    "Yusril",
    "Arzaqi",
    "Bimo",
    "Adam",
  }

  for index, name:= range names {
    fmt.Println("Index", index, "=", name)
  }
}
```

---

**Menggunakan Map**

```go
func main(){
  person := map[string]string{
		"name":  "Yusril",
		"kelas": "TKJ2",
		"nim":   "185512",
	}

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}
```

## Break & continue

- `break` & `continue` adalah kata kunci yang bisa digunakan dalam perulangan.
- `break` digunakan untuk menghentikan seluruh perulangan.
- `continue` adalah digunakan untuk menghentikan perulangan yang berjalan, dan langsung melanjutkan ke perulangan selanjutnya.nya

### Kode : Break

```go
func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke ", i)
	}
}
```

**Hasil**

```
Perulangan ke  0
Perulangan ke  1
Perulangan ke  2
Perulangan ke  3
Perulangan ke  4
```

### Kode : Continue

```go
func main() {
  for i := 0; i < 10; i++ {
    if i % 1 == 0 {
      continue
    }
    fmt.Println("Perulangan ke", i)
  }
}
```

## Function

- Sebelumnya kita sudah mengenal sebuah function yang wajib dibuat agar program kita bisa berjalan, yaitu function main.
- Function adalah sebuah blok kode yang sengaja dibuat dalam program agar bisa digunakan berulang-ulang.
- Cara membuat function sangat sederhana, hanya dengan menggunakan kata kunci `func` lalu diikuti dengan nama function nya dan blok kode isi function nya.
- Setelah membuat function, kita bisa mengeksekusi function tersebut dengan memanggilnya menggunakan kata kunci nama functionnya diikuti tanda kurung buka `(`, kurung tutup `)`

### Kode : Function

```go
func sayHello() {
  fmt.Println("Hello ")
}

func main() {
  sayHello()
}
```

## Function Parameter

- Saat membuat function, kadang-kadang kita membutuhkan data dari luar, atau kita sebut parameter.
- Kita bisa menambahkan parameter di function, bisa lebih dari satu.
- Parameter tidaklah wajib, jadi bisa membuat function tanpa parameter seperti sebelumnya yang sudah kita buat.
- Namun jika kita menambahkan parameter di function, maka ketika memanggil function tersebut, kita wajib memasukkan data ke parameter.

### Kode : Function Parameter

```go
func sayHelloTo(firstName string, lastName string)  {
  fmt.Println("Hello", firstName, lastName)
}

func main() {
  sayHelloTo("Yusril", "Arzaqi")
}
```

## Function Return Value

- Function bisa mengembalikan data.
- Untuk memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data kembalikan dari function tersebut.
- Jika function tersebut kita deklarasikan dengan tipe data pengembalikan, maka wabjib di dalam function nya kita harus mengembalikan data.
- Untuk mengembalikan data dari function, kita bisa menggunakan kata kunci `return`, diikuti dengan datanya.

### Kode : Function Return Value

```go
func sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}
func getHello(name string) string {
	if name == "" {
		return "Hello Anda Siapa"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Yusril")
	fmt.Println(result)

	result = getHello("")
	fmt.Println(result)

	fmt.Println(sum([]int{10, 20, 30, 40}))
}
```

## Returning Multiple Value

- Function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value.
- Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data `return` value di function.

### Menghiraukan Return Value

- Multiple `return` value wajib ditangkap semua value nya.
- Jika kita ingin menghiraukan `return` value tersebut, kita bisa menggunakan tanda `_` (garis bawah).

### Kode : Menghiraukan Return Value

```go
func main() {
  firstName, _ := getFullName()
  fmt.Println(firstName)
}
```

## Named Return Value

- Biasanya saat kita memberitahu bahwa sebuah function mengembalikan value, maka kita hanya mendeklarasikan tipe data `return` value di function.
- Namun kita juga bisa membuat variable secara langsung di tipe data `return` function nya.

### Kode : Named Return Value

```go
func getComplateName() (firstName string, lastName string){
  firstName = "Yusril"
  lastName = "Arzaqi"
  return firstName, lastName
}
func main()  {
  firstName, lastName := getComplateName()
  fmt.Println(firstName, lastName)
}
```

## Variadic Function

- Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varags.
- Varags artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array.
- Apa bedanya dengan parameter biasa dengan tipe data Array ?
  - Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function.
  - Jika parameter menggunakan varags, kita bisa langsung mengirim datanya, jika lebih dari satu, cukup gunakan tanda koma.

### Kode : Variadic Function

```go
func sumAll(numbers ...int) int {
  total := 0
  for _, number := range numbers {
    total += number;
  }
  return total;
}

func main() {
  total := sumAll(10, 20, 30, 40, 50, 60, 70, 80, 90)
  fmt.Println(total)
}
```

### Slice Parameter

- Kadang ada kasus dimana kita menggunakan Variadic Function, namun memiliki variable berupa slice.
- Kita bisa menjadikan slice sebagai varag parameter.

### Kode : Slice Parameter

```go
func main() {
  numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
  result = sumAll(numbers...)
  fmt.Println(result)
}
```

## Function Value

- Function adalah first class citizen.
- Function juga merupakan tipe data, dan bisa disimpan di dalam variable.

### Kode : Function Value

```go
func getGoodBye(name string) string {
  return "Good Bye " + name
}

func main() {
  goodbye := getGoodBye;
  fmt.Println(goodbye("Yusril"))
  fmt.Println(getGoodBye("Arzaqi"))
}
```

## Function as Parameter

- Function tidak hanya bisa kita simpan di dalam variable sebagai value.
- Namun juga bisa kita gunakan sebagai parameter untuk function lain.

### Kode : Function as Parameter

```go
func sayHelloWithFilter(name string, filter func(string)string ) {
  fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
  if name == "Anjing" {
    return "..."
  } else {
    return name
  }
}

func main() {
  sayHelloWithFilter("Anjing", spamFilter)
  sayHelloWithFilter("Yusril", spamFilter)
}
```

### Function Type Declarations

- Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter.
- Type Declarations juga bisa digunakan untuk membuat alias function, sehingga akan memepermudah kita menggunakan function sebagai parameter.

## Anonymous Function

- Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut.
- Namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membuat function terlebih dahulu.
- Hal tersebut dinamakan anonymous function, atau function tanpa nama.

### Kode : Anonymous Function

```go
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
  if blacklist(name) {
    fmt.Println("You Are Blocked", name)
  } else {
    fmt.Println("Welcome", name)
  }
}

func main() {
  blacklist := func(name string) bool {
    return name == "anjing"
  }
  registerUser("Yusril", blacklist);
  registerUser("anjing", blacklist);
  registerUser("Dimas", func(name string) bool {
    return name == "anjing"
  })
}
```

## Recursive Function

- Recursive function adalah function yang memanggil function dirinya sendiri.
- Kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive function.
- Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah **Faktorical**

### Kode : Factorical For Loop

```go
func factoricalLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(factoricalLoop(5))
}
```

### Kode : Faktorical For Recursive

```go
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * (factorialRecursive(value - 1))
	}
}

func main() {
	fmt.Println("faktorical menggunakan loop", factoricalLoop(5))
	fmt.Println("faktorical menggunakan Recursive", factorialRecursive(5))
	fmt.Println("faktorical menggunakan manual", 1*2*3*4*5)
}
```

## Closures

- Closures adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama.
- Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi.

### Kode : Closures

```go
func main() {
  name := "Yusril"
  counter := 0

  increment := func() {
    const name = "Dimas"
    fmt.Println("Increment")
    counter++; // counter akan bisa di akses
  }

  increment()
  increment()
  fmt.Println(counter)
  fmt.Println(name)
}
```

## Defer, Painc dan Revocer

### Apa itu Defer

- Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dari eksekusi.
- Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi.

### Kode : Defer

```go
func logging() {
  fmt.Println("Selesai memanggil function")
}

func runApplication() {
  defer logging()
  fmt.Println("Run Application")
}

func main() {
  runApplication()
}
```

### Panic

- Panic function adalah function yang bisa kita gunakan untuk menghentikan program.
- Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan.
- Saat `panic` function biasanya dipanggil, program akan terhenti, namun `defer` function tetap akan dieksekusi.

### Recover

- Recover adalah function yang bisa kita gunakan untuk menangkap data `panic`.
- Dengan `recover` process `panic` akan terhenti, sehingga program akan tetap berjalan.

#### Kode : Recover Salah

```go
func endApp() {
  fmt.Prinln("Selesai memanggil function")
}
func runApp(error bool) {
  defer endApp()
  if error {
    panic("ERROR")
  }
  result := recover()
  fmt.Println("Terjadi Error", result)
  rmt.Println("Aplikasi berjalan")
}

func main() {
  runApp(true)
}
```

#### Kode : Recover Benar

```go
func endApp() {
  fmt.Println("Selesai memanggil function")
  result := recover()
  fmt.Println("Terjadi ERROR", result)
}

func runApp(error bool) {
  defer endApp();
  if error {
    panic("Error")
  }
  fmt.Println("Aplikasi berjalan")
}

func main() {
  runApp(ture)
}
```

## Komentar

- Komentar terbaik pada kode adalah kode itu sendiri.
- Saat membuat kode, kita perlu membuat kode semudah mungkin untuk dibaca.
- Namun kadang juga kita butuh menambahkan komentar dikode kita.

### Kode : Komentar

```go
package main

/**
This is multi line comment
You can add a lot of comments here
**/

func main() {
  // this is single line comment
}
```

## Struct

- `struct` adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan.
- `struct` biasanya representasi data dalam program aplikasi yang kita buat.
- Data di `struct` disimpan dalam filed.
- Sederhananya `struct` adalah kumplan dari field.

### Kode Struct

```go
type Customer struct {
	Name, Address string
	Age           int
}
```

### Membuat Data Struct

- `struct` adalah template data atau prototype data.
- `struct` tidak bisa langsung digunakan.
- Namun kita bisa membuat data/object dari struct yang telah kita buat.

### Kode : Membuat Data Struct

```go
func main() {
  var yusril Customer
  yusril.Name = "Yusril Arzaqi"
  yusril.Address = "Indonesia"
  yusril.Age = 18

  fmt.Println(yusril)
  fmt.Println(yusril.Name)
  fmt.Println(yusril.Address)
  fmt.Println(yusril.Age)
}
```

### Struct Literals

- Sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untukmembuat data dari struct.

### Kode : Struct Literals

```go
bimo := Customer {
  Name: "Bimo Alamsyah",
  Address: "Indonesia",
  Age: 16,
}

fmt.Println(bimo)
fmt.Println(bimo.Name)
fmt.Println(bimo.Address)
fmt.Println(bimo.Age)

adam := Customer{"Adam Saputra", "Indonesia", 17}

fmt.Println(adam)
fmt.Println(adam.Name)
fmt.Println(adam.Address)
fmt.Println(adam.Age)
```

## Struct Method

- `struct` adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function.
- Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function.
- Method adalah function.

### Kode : Struct Method

```go
func (customer Customer) sayHi(name string) {
  fmt.Println("Hello,", name, "My Name is", customer.Name, "salam kenal.")
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, My Neme is", customer.Name)
}

func main() {
  yusril := Customer{Name: "Yusril Arzaqi"}
  yusril.sayHello() // Hello, My Neme is Yusril Arzaqi
  yusril.sayHi("Bimo") // Hello, Bimo My Name is Yusril Arzaqi salam kenal.
}
```

## Interface

- `interface` adalah tipe data Abstract, dia tidak memiliki implementasi langsung.
- Sebuah `interface` berisikan definisi-definisi method.
- Biasanya `interface` digunakan sebagai kontrak.

### Kode : Interface

```go
type HasName interface {
  GetName() string
}

func SayHello(hasName HasName) {
  fmt.Println("Hello", hasName.GetName())
}
```

### Implementasi Interface

- Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut.
- Sehingga kita tidak perlu mengimplementasikan interface secara manual.
- Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika memubat interface, kita harus menyebutkan secara eksplisit akan menggunakan intreface mana.

### Kode : Implementasi Interface

```go
type Person struct {
  Name string
}

func (person Person) GetName() string {
  return person.Name
}


func main() {
  yusril := Person{Name: "Yusril Arzaqi"}
  myName := yusril.GetName()
  SayHello(yusril)
  fmt.Println(myName)
}
```

## Interface Kosong

- Go-Lang bukanlah bahasa pemrograman yang berorientasi objek.
- Biasanya dalam pemrograman berorientasi object. ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut.
- Contoh di Java ada `java.lang.Object`.
- Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong.
- Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya.

### Penggunaan Interface Kosong di Go-Lang

- Ada banyak contoh penggunaan interface kosong di Go-Lang, seperti :
  - `fmt.Println(a ...interface{})`.
  - `panic(v interface{})`.
  - `recover() interface{}`.
  - dan lain-lain.

### Kode : Interface Kosong

```go
func Ups() interface{} {
  return "Ups"
}

func main() {
  kosong := Ups()
  fmt.Println(kosong)
}
```

## Nil

- Nil Biasanya didalam bahasa pemrograman lain, object yang belum diinialisasi maka secara otomatis nilainya adalah `null` atau `nil`.
- Berbeda dengan Go-Lang, di Go-lang saat kita buat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan default value nya.
- Namun di Go-Lang ada data `nil`, yaitu data kosong.
- Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti `interface`, `func`, `map`, `slice`, `pointer`, dan `channel`.

### Kode : Nil

```go
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
  yusril := NewMap("Yusril Arzaqi")

  if yusril == nil {
    fmt.Println("Data Kosong")
  } else {
    fmt.Println(yusril)
  }
}
```

## Error Interface

- Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interfacenya adalah error.

```go
// The error built-in interface
type error interface {
  Error() string
}
```

## Type Assertions

- Type Assertion merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan.
- Fitur ini sering sekali digunakan ketika bertemu dengan data interface kosong.

## Kode : Type Assertions

```go
func random() any {
  return "OK"
}

func main() {
  result := random()
  resultString := result.(string)
  fmt.Println(resultString)

  resultInt := result.(int) // panic
  fmt.Println(resultInt)
}
```

### Type Assertions Menggunakan Switch

- Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic diaplikasi kita.
- Jika panic dan tidak ter-recover, maka otomatis program kita akan mati.
- Agar lebih aman, sebaliknya kita menggunakan `switch` expression untuk melakukan text assertions.

### Kode : Type Assertions Menggunakan Switch

```go
func random() interface{} {
  return "OK"
}

func main() {
  result := random()
  switch value := result.(type) {
  case string:
    fmt.Println("String", value)
  case int:
    fmt.Println("Int", value)
  default:
    fmt.Println("Unknown Type : ", value)
  }
}
```

## Pointer

- Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada.
- Sederhananya, dengan kemampuan pointer, kita bisa membuat _pass by reference_

### Pass by Value

- Secara default di Go-Lang semua variable itu di passing by value, bukan by reference.
- Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi value nya.

### Kode : Pass by Value

```go
type Address struct {
  City, Province, Country string
}

func main() {
  address1 := Address{
    "Semarang",
    "Jawa Tengah",
    "Indonesia",
  }

  address2 := address1
  address2.City = "Demak"

  fmt.Println(address1) // address1 tidak berubah
  fmt.Println(address2)
}
```

### Penjelasan Detail Passing by Value

![Penjelasan Detail Passing by Value](./img/Penjelasan-Detail-Pass-by-Value.png)

### Operator & (pointer)

- Untuk memubat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan opertaor `&` diikuti dengan nama variablenya

### Kode : Operator & (pointer)

```go
func main() {
  address1 := Address{
    "Semarang",
    "Jawa Tengah",
    "Indonesia",
  }

  address2 := &address1
  address2.City = "Demak"

  fmt.Println(address1) // address1 berubah
  fmt.Println(address2)
}
```

### Operator \* (pointer golang)

- Saat kita merubah variable pointer, maka yang berubah hanya variable tersebut.
- Semua variable yang mengacu ke data yang sama tidak akan berubah.
- Jika kita ingin mengubah variable yang menhacu ke data tersebut, kita bisa menggunakan operator \*

### Kode : Operator \* (pointer golang)

Tidak Menggunakan pointer

```go
func main() {
  address1 := Address{
    "Semarang",
    "Jawa Tengah",
    "Indonesia",
  }

  address2 := &address1
  address2.City = "Demak"

  address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}

  fmt.Println(address0) // address1 akan berubah
  fmt.Println(address1)
}
```

---

Menggunakan pointer

```go
func main() {
  address0 := Address{
    "Semarang",
    "Jawa Tengah",
    "Indonesia",
  }

  address1 := &address1
  address1.City = "Demak"
  var address2 *Address = &address1

  fmt.Println(address1)

  *address1 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

  fmt.Println(address0) // address1 akan berubah
  fmt.Println(address1) // berubah menyerupai address1
  fmt.Println(address2) // menjadi address1
}
```

### Pass by Reference dengan Pointer

![Pass by Reference dengan Pointer](./img/Pass-by-Refrence-dengan-Pointer.png)

### Tanpa Operator \*

![Tanpa Operator](./img/Tanpa-Operator*.png)

### Dengan Operator \*

![Dengan Operator](./img/dengan-operator*.png)

### Function `new`

- Sebelumnya untuk membuat pointer dengan menggunakan operator \&.
- Go-Lang juga memiliki function `new` yang biasa digunakan untuk membuat pointer.
- Namun function `new` hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal.

### Kode : Function `new`

```go
func main() {
  alamat1 := new(Address)
  alamat2 := alamat1

  alamat2.Country = "Indonesia"

  fmt.Println(alamat1) // alamat 1 akan berubah
  fmt.Println(alamat2)
}
```

## Pointer di Function

- Saat kita membuat parameter di function, secara default adalah _pass by value_, artinya datanya akan di copy lalu dikirim ke function tersebut.
- Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah.
- Hal ini membuat variable menjadi akan, karena tidak akan bisa diubah.
- Namun kadang kita ingin membuat akan, karena tidak akan bisa diubah.
- Untuk melakukan ini, kita juga bisa menggunakan pointer di function.
- Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator \* di parameternya.

### Kode : Pointer di Function

```go
type Address struct {
  City, Province, Country string
}

func ChangeAddressToIndonesia(address Address) {
  address.Country = "Indonesia"
}

func main() {
  address := Address{
    "Semarang",
    "JawaTengah",
    "",
  }

  fmt.Println(address)
  ChangeAddressToIndonesia(*address)
  fmt.Println(address) // tidak akan berubah
}
```

---

```go
func ChangeAddressToIndonesia(address *Address) {
  address.Country = "Indonesia"
}

func main() {
  address := Address{
    "Semarang",
    "JawaTengah",
    "",
  }

  fmt.Println(address)
  ChangeAddressToIndonesia(&address)
  fmt.Println(address) // akan berubah
}
```

## Pointer di Method

- Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah _pass by value_.
- Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil mehod

### Kode : Pointer di Method

```go
func (man Man) Married() {
  man.Name = "Mr, " + man.Name
  println(man.Name)
}

func main() {
  yusril := Man{"Yusril"}
  fmt.Println(yusril)
  yusril.Married()
  fmt.Println(yusril)
}
```

---

```go
func (man *Man) Married() {
  return "Mr. " + man.Name
}

func main() {
  yusril := Man{"Yusril"}
  fmt.Println(yusril)
  yusril.Married()
  fmt.Println(yusril)
}
```

## GOPATH

- Sebelumnya saat kita membuat project aplikasi Go-Lang kita belum membahas tentang GOPATH.
- GOPATH adalah environment variable yang berisikan lokasi tempat kita menyimpan project dan library Go-Lang.
- GOPATH wajib di buat ketika kita membuat aplikasi lebih dari satu file atau butuh menggunakan library dari luar.

### Export

```bash
export GOPATH="/home/yusril/Developments/GoLang"
```

## Package Import

### Package

- Package adalah tempat yang bisa digunakan untuk mengorganisir kode program yang kita buat di Go-Lang.
- Dengan menggunakan package, kita bisa merapikan kode program yang kita buat.
- package sendiri sebenarnya hanya direktori folder di sistem operasi kita.

```go
package helper

func SayHello(name string) string {
  return "Hello" + name
}
```

### import

- Secara standar, file Go-Lang hanya bisa mengakses file Go-Lang lainnya yang berada dalam package yang sama.
- Jika kita ingin mengakses file Go-Lang yang berada diluar package, maka kita bisa menggunakan imporjt.

### Kode : Import

```go
package main

import  "belajar-golang-dasar/helper"

func main() {
  helper.SayHello("Yusr")
}
```

## Access Modifier

- Di bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadap suatu function dan variable.
- Di Go-Lang, untuk menentukan access modifier, cukup dengan nama function atau variable.
- Jika namanya diawali dengan, huruf besar, maka artinya bisa diakses dari package lain, jika dimulai dengan huruf kecil, artinya tidak bisa diakses dari package lain.

### Kode : Access Modifier

```go
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
```

## Package Initialization

- Saat kita membuat package, kita bisa membuat sebuah function yang akan diakses ketika package kita diakses.
- Ini sanggat cocok ketika contohnya, jika package kita berisi function-function untuk berkomunikasi dengan database, kita membuat function inisialisasi untuk membuat koneksi ke database.
- Untuk membuat function yang diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama init.

### Kode : Package Initialization

```go
package database

var connection string

func init() {
  connection = "MySQL"
}

func GetDatabase() string {
  return connection
}
```

---

```go
package main

import (
	"belajar-golang-dasar/database"
	"fmt"
)

func main() {
  result := database.GetDatabase()
  fmt.Println(result)
}
```

### Blank Identifier

- Kadang kita hanya ingin menjalankan `init` function di package tanpa harus mengeksekusi salah satu function yang ada di package.
- Secara defualt, Go-Lang akan komplen ketika ada package yang di import namun tidak digunakan.
- Untuk menangani hal tersebut, kita bisa menggunakan blank identifier `_` sebelum nama package ketika melakukan import.

## Package OS

- Go-Lang telah menyediakan banyak sekali package bawaan, saah satunya adalah package os.
- Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara indpenden (bisa digunakan disemua sistem operasi)
- [OS](https://golang.org/pkg/os)

### Kode Package os

```go
package main

import (
  "fmt"
  "os"
)

func main() {
  args := os.Args
  fmt.Println(args)
}
```

## Package flag

- Package flag berisikan fungsionalitas untuk memparsing command line argument.
- [flag](https://golang.org/pkg/flag)

## Kode : Package flag

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	///////////////////(nama flag), (default), (description)
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
  var number = flag.Int("number", 0, "Put your number")

	flag.Parse()

  fmt.Println("host : ",*host)
  fmt.Println("Username : ",*username)
  fmt.Println("Password : ",*password)
  fmt.Println("number : ", *number)
}
```

## Package Strings

- Package strings adalah package yang berisika function-function untuk memanipulasi tipe data Strings.
- Ada baynak sekali function yang bisa kita gunakan.
- [strings](https://golang.org/pkg/strings/)

### Beberapa Function di Package strings

|                Fucntion                |                     Kegunaan                     |
| :------------------------------------: | :----------------------------------------------: |
|     `strings.Trim(string, cutset)`     |       Memotong cutset di awal akhir string       |
|       `strings.ToLower(string)`        | Membuat semua karakter string menjadi lower case |
|       `strings.ToUpper(string)`        | Membuat semua karakter string menjadi upper case |
|   `strings.Split(string, separator)`   |      Memotong string berdasarkan separator       |
|   `strings.Contains(string, search)`   |  Mengecek apakah string mengandung string lain   |
| `strings.ReplaceAll(string, from, to)` |      Mengubah semua string dari from ke to       |

### Kode : Package Strings

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
  fmt.Println(strings.Contains("Yusril Arzaqi Mantap","Yusril"))
  fmt.Println(strings.Split("Yusril Arzaqi Mantap", " "))
  fmt.Println(strings.ToLower("Yusril Arzaqi"))
  fmt.Println(strings.ToUpper("Yusril Arzaqi"))
  fmt.Println(strings.Trim("                  Yusril Arzaqi           ", " "))
  fmt.Print(strings.ReplaceAll("Yusril Yusril Yusril Yusril Yusril", "Yusril", "Arzaqi"))
}
```

## Package strconv (string conversion)

- sebaliknya kita sudah belajar cara konversi tipe data, misal dari `int32` ke `int64`.
- Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda ? Misal dari int ke string, atau sebaliknya.
- Hal tersebut bisa kita lakukan dengan bantukan package `strconv` (string conversion)
- [strconv](https://golang.org/pkg/strconv)

### Beberapa Function di Package `strconv`

|            Function             |            Kegunaan            |
| :-----------------------------: | :----------------------------: |
|   `strconv.parseBool(string)`   |  Mengubah `string` ke `bool`   |
|  `strconv.parseFloat(string)`   |   Mengubah `string` `float`    |
|   `strconv.parseInt(string)`    |  Mengubah `string` ke `int64`  |
|   `strconv.FormatBool(bool)`    |  Mengubah `bool` ke `string`   |
| `strconv.FormatFloat(int, ...)` | Mengubah `float64` ke `string` |
|  `strconv.FormatInt(int, ...)`  |  Mengubah `int64` ke `string`  |

### Kode : Package `strconv`

```go
func main() {
  boolean, err := strconv.ParseBool("true")

  if err == nil {
    fmt.Println(boolean)
  } else {
    fmt.Println("Error", err.Error())
  }
}
```

## Package Math

- Package math merupakan package yang berisikan constant dan fungsi Matematika.
- [Math](https://golang.org/pkg/math/)

### Beberapa Function di Package math

|           Function           |                                  Kegunaan                                  |
| :--------------------------: | :------------------------------------------------------------------------: |
|    `math.Round(float64)`     | Membulatkan `float64` keatas atau kebawah, sesuai dengan yang paling dekat |
|    `math.Floor(float64)`     |                       Membulatkan `float64` kebawah                        |
|     `math.Ceil(float64)`     |                        Membulatkan `float64` keatas                        |
| `math.Max(float64, float64)` |                 Mengembalikan nilai `float64` paling besar                 |
| `math.Min(float64, float64)` |                 Mengembalikan nilai `float64` paling kecil                 |

### Kode : Package Math

```go
package main

import (
	"fmt"
	"math"
)

func main() {
  fmt.Println(math.Round(1.7))
  fmt.Println(math.Round(1.4))
  fmt.Println(math.Floor(1.7))
  fmt.Println(math.Ceil(1.4))

  fmt.Println(math.Max(10, 20))
  fmt.Println(math.Min(10, 20))
}
```

## Package container / list

- Pcakage container / list adalah implementasi struktur data double linked list di Go-Lang.
- [container list](https://golang.org/pkg/container/list)

### Struktur Data Double Linked List

![Double Linked](./img/double-linked.png)

### Kode : Package container / list

```go
package main

import (
	"container/list"
	"fmt"
)

func main() {
  data := list.New()
  data.PushBack("Yusril")
  data.PushBack("Arzaqi")
  data.PushBack("Bimo")
  data.PushBack("Alamsyah")

  for e := data.Front(); e != nil; e = e.Next() {
    if e.Value == "Yusril" {
      fmt.Println("Horay")
    }
    fmt.Println(e.Value)
  }
}
```

## Package container / ring

- Package container/ring adalah implementasi struktur data ciruclar list.
- Circular list adalah struktur data ring, dimana diakhir element akan kembali ke element awal (Head).
- [container/ring](https://golang.org/pkg/container/ring/)

### Struktur Data Circular List

![Struktur Data Circular List](./img/data-circular-list.png)

### Kode : Package container / ring

```go
package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
```

## Package sort

- Package sort adalah package yang berisikan utilitas untuk proses pengurutan.
- Agar data kita bisa diurutkan, haruus mengimplementasikan kontrak di interface `sort.Interface`.
- [sort](https://golang.org/pkg/sort)

### `sort.interface`

```go
type Interface interface {
  Len() int
  Less(i, j int) bool
  Swap(i, j int)
}
```

### Kode : Package `sort`

```go
package main

import (
	"fmt"
	"sort"
)

type User struct {
  Name string
  Age int
}

type UserSlide []User

func (userSlice UserSlide) Len() int {
  return len(userSlice)
}

func (userSlice UserSlide) Less(i, j int) bool {
  return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlide)  Swap(i, j int) {
  userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
  users := UserSlide {
    {"Yusril", 18},
    {"Arzaqi", 18},
    {"Bimo", 17},
    {"Alamsyah", 17},
    {"Adam", 19},
  }

  fmt.Printf("users: %v\n", users)
  sort.Sort(users)
  fmt.Printf("users: %v\n", users)
}
```

## Package time

- Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang.
- [time](https://golang.org/pkg/time)

### Beberapa Function di Package time

|           Function           |              Kegunaan               |
| :--------------------------: | :---------------------------------: |
|         `time.Now()`         |  Untuk mendapatkan waktu saat ini.  |
|        `time.Date()`         |        Untuk membuat waktu.         |
| `time.Parse(layout, string)` | Untuk memparsing waktu dari string. |

### Kode : Package time

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now.Local())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Unix())

	utc := time.Date(2003, time.July, 23, 23, 59, 59, 0, time.UTC)
	fmt.Println(utc.Local())

	parse, err := time.Parse(time.RFC3339, "2003-07-23T23:59:59Z")
	if err == nil {
		fmt.Println(parse)

	} else {
		fmt.Println(err.Error())
	}
}
```

## Package reflect

- Dalam bahasa pemrograman, biasanya ada futir **Reflection**, dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan.
- Hal ini bisa dilakukan di Go-Lang dengan menggunakan ackage reflect.
- Fitur ini mungkin tidak bisa dibahas secara lengkap dalam satu video, Anda bisa eksplorasi package `reflec` ini secara otodidak.
- Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan.
- [reflect](https://golang.org/pkg/reflect)

### Kode : Package reflect

```go
package main

import (
  "fmt"
  "reflect"
)

type Sample struct {
  // Name string
  Name string `required:"true" max:"9"`
}

func main() {
  sample := Sample{"Yusril"}
  sampleType := reflect.TypeOf(sample)
  structField := sampleType.Field(-1)
  // required := structField.Tag.Get("required")

  fmt.Printf("sample: %v\n", sample)
  fmt.Println("sampleType", sampleType)
  fmt.Println(structField.Type)
  fmt.Println(sampleType.Field(-1).Tag.Get("required"))
  fmt.Println(sampleType.Field(-1).Tag.Get("max"))
  fmt.Println(IsValid(sample))
}
```

### Kode : Validation

```go
package main

import "reflect"

func IsValid(data any) bool {
  t := reflect.TypeOf(data)

  for i := 0; i < t.NumField(); i++ {
    field := t.Field(i)
    if field.Tag.Get("required") == "true" {
      return reflect.ValueOf(data).Field(i).Interface() != ""
    }
  }

  return true
}
```

## Package regexp

- Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular exporession.
- Regular expression di Go-Lang menggunakan lirary C yang dibuat Google bernama RE2.
- [RE2 Symtax](https://github.com/google/re2/wiki/syntax)
- [regex](https://golang.org/pkg/regexp)

### Beberapa Function di package regexp

|              Function               |                       Kegunaan                        |
| :---------------------------------: | :---------------------------------------------------: |
|    `regexp.MustCompile(string)`     |                    Membuat Regexp                     |
|  `Regexp.MatchString(string) bool`  |      Mengecek apakah Regexp match dengan string       |
| `Regexp.FindAllString(string, max)` | Mencari string yang match dengan maximum jumlah hasil |

### Kode : Package Regexp

```go
package main

import (
	"fmt"
	"regexp"
)


func main() {
	regex := regexp.MustCompile(`e([a-z])`) 

	fmt.Println(regex.MatchString("eko")) // true
	fmt.Println(regex.MatchString("edo")) // true
	fmt.Println(regex.MatchString("eKo")) // false

  search := regex.FindAllString("eko edo egi ego e1o eto", 10)

  for _, v := range search {
    fmt.Println(v)
  }
}
```

## Materi selanjutnya ?

- Go-Lang Unit test
- Go-Lang Modules
- Go-Lang concurrency
- Go-Lang Database
- Go-Lang Web
