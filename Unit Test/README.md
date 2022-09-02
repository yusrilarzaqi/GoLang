# Unit Test

## Sebelum Belajar

- Go-Lang Dasar
- Go-Lang Modules

## Agenda

- Pengenalan Software Testing.
- testing Package.
- Unit Test.
- Assertion.
- Mock.
- Benchmark.

## Pengenalan Software Testing

- Software testing adalah salah atu disiplin ilmu dalam Software engineering.
- Tujuan utama dari software testing adalah memastikan kualitas kode dan aplikasi kita baik.
- Ilmu untuk software testing sangatlah luas, pada materi ini kita hanya akan fokus ke unit testing.

### Test Pyramid

![Test Pyramid](./img/test-pyramid.png)

### Contoh High Level Architecture Aplikasi

![Contoh High Level Architecture Aplikasi](./img/contoh-high-level-architecture-aplikasi.png)

### UI Test / End to End Test

![UI Test / End to End Test](./img/UI-Test.png)

### Service Test / Integration Test

![Service Test / Integration Test](./img/Service-Test.png)

### Contoh Internal Architecture Aplikasi

![Contoh Internal Architecture Aplikasi](./img/contoh-internal-architecture-aplikasi.png)

### Unit Test diagram

![Unit Test](./img/Unit-Test.png)

### Pengertian Unit Test

- Unit test akan fokus menguji bagian kode program terkecil, biasanya menguji sebuah method.
- Unit test biasanya dibuat kecil dan cepat, oleh karena itu biasanya kadang kode unit test lebih banyak dari kode program aslinya, karena semua skenario pengujian akan dicoba di unit test.
- Unit test bisa digunakan sebagai cara untuk mengigkatkan kualitas

## Pengenalan Testing Package

- Di bahasa pemrograman lain, biasanya untuk implementasi unit test, kita butuh library atau framework khusus untuk unit test.
- Berbeda dengan Go-Lang, di Go-Lang sudah untuk unit test sudah disediakan sebuah package khusus bernama `testing`.
- Selain itu untuk menjalankan unit test, di Go-Lang juga sudah disediakan perintahnya.
- Hal ini membuat implementasi unit testing di golang sangat mudah dibanding dengan bahasa pemrograman yang lain.
- [testing](https://golang.org/pkg/testing)

### `testing.T`

- Go-Lang menyediakan sebuah struct bernama `testing.T`.
- Struct ini digunakan untuk unit test di Go-Lang.

### `testing.M`

- `testing.M` adalah struct yang disediakan Go-Lang mengatur life cycle testing.
- Materi ini nanti akan kita bahas di cahpter main.

### `testing.B`

- `testing.B` adalah struct yang disediakan Go-Lang untuk melakukan benchmarking.
- Di Go-Lang untuk melakukan benchmark (mengukur kecepatan kode program) pun sudah disediakan, sehingga kita tidak perlu menggunakan library atau framework tambahan.

## Membuat unit test

### Kode : Hello World Function

```go
package helper

func HelloWorld(name string) string {
  return "Hello " + name
}
```

### Aturan File Test

- Go-Lang memiliki aturan cara membuat file khusus untuk unit test.
- Untuk membuat file unit test, kita harus menggunakan akhiran `_test`.
- Jadi kita misal kita membuat file `hello_world.go`, artinya untuk membuat unit testnya, kita harus memubat file `hello_world_test.go`.

### Aturan Function Unit Test

- Selain aturan nama file, di Go-Lang juga sudah diatur untuk nama function unit test.
- Nama function untuk unit test harus diawali dengan nama Test.
- Misal jika ingin mengetest function HelloWorld, maka kita akan membuat function unit test dengan nama `TestHelloWorld`.
- Tidak ada aturan untuk nama belakang function unit test harus sama dengan nama function yang akan di test, yang penting adalah harus diawali dengan kata `Test`.
- Selanjutnya harus memiliki parameter `t *testing.T` dan tidak mengembalikan return value.

### Kode : Hello World Unit Test

```go
package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
  result := HelloWorld("Yusril")

  if result != "Hello Yusril" {
    // unit test failed
    panic("Result is not Hello Yusril")
  }

  fmt.Println(result)
}
```

### Menjalankan Unit Test

- Untuk menjalankan unit test kita bisa menggunakan perintah :
  - `go test`
- Jika kita ingin lihat detail function test apa saja yang sudah di running, kita bisa gunakan perintah :
  - `go test -v`.
- Dan jika kita ingin memilih function unit test mana yang ingin dirunnung, kita bisa gunakan perintah :
  - `go test -v -run TestNamaFunction`

### Menjalankan Semua Unit Test

- Jika kita ingin menjalankan semua unit test dari top folder modulenya, kita bisa gunakan perintah :
  - `go test ./..`

## Menggagalkan Unit Test

- Menggagalkan unit test menggunakan `panic` bukanlah hal yang bagus.
- Go-Lang sendiri sudah menyediakan cara untuk menggagalkan unit test menggunakan `testing.T`.
- Terdapat function `Fail()`, `FailNow()`, `Error()` dan `Fatal()` jika kita ingin menggagalkan unit test.

### `t.Fail()` dan `t.FailNow()`

- Terdapat dua function untuk menggagalkan unit test, yaitu `Fail()` dan `FailNow()`. LAntas apa bedanya ?
- `Fail()` akan menggagalkan unit test, namun tetap melanjutakan eksekusi unit test. Namun diakhir ketika selesai, maka unit test tersebut dianggap gagal.
- `FailNow()` akan menggaalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test.

---

```go
func TestHelloWorldYusril(t *testing.T) {
	result := HelloWorld("Yusril")

	if result != "Hello Yusril" {
		// unit test failed
		t.Fail()
	}

	fmt.Println("TestHelloWorldYusril done")
}
```

---

```go
func TestHelloWorldArzaqi(t *testing.T) {
	result := HelloWorld("Arzaqi")

	if result != "Hello Arzaqi" {
		// unit test failed
		t.FailNow()
	}

	fmt.Println("TestHelloWorldArzaqi done")
}
```

### `t.Error(args...)` dan `t.Fatal(args...)`

- Selain `Fail()` dan `FailNow()`, ada juga `Error()` dan `Fatal()`.
- `Error()` function lebih seperti melakukan log `print` error, namun setelah melakukan log error, dia akan secara otomatis memanggil function `Fail()`, sehingga mengakibatkan unit test dianggap gagal.
- Namun karena hanya memanggil `Fail()`, artinya eksekusi unit test akan tetap berjalan sampai selesai.
- `Fatal()` mirip dengan `Error()`, hanya saja, setelah melakukan log error, dia akan memanggil `FailNow()`, sehingga mengakibatkan eksekusi unit test berhenti.

---

```go
func TestHelloWorldBimo(t *testing.T)  {
  result := HelloWorld("Bimo")

  if result != "Hello Bimo" {
    t.Error("Harusnya Hello Bimo")
  }

  fmt.Println("Dieksekusi")
}
```

---

```go
func TestHelloWorldAlamsyah(t *testing.T) {
  result := HelloWorld("Alamsyah")

  if result != "Hello Alamsyah" {
    t.Fatal("Harusnya Hello Alamsyah")
  }

  fmt.Println("Tidak Dieksekusi")
}
```

## Assertion

- Melakukan pengecekan di unit test secara menual menggunakan `if`, `else` sangatlah menyebalkan.
- Apalagi jika result yang harus di cek itu banyak.
- Oleh karena itu, sangat disarankan untuk menggunakan Assertion untuk melakukan pengecekan.
- Sayangnya, Go-Lang tidak menyediakan package untuk assertion, sehingga kita butuh menambahkan library untuk melakukan assertion ini.

### Testify

- Salah satu library assertion yang paling populer di Go-Lang adalah Testify.
- Kita bisa menggunakan library ini untuk melakukan assertion terhadap result data di unit test.
- [testify](https://github.com/stretchr/testify)
- Kita bisa menambahkan di Go module kita :
  - `go get github.com/stretchr/testify`

### `assert` vs `require`

- Testify menyediakan dua package assertion, yaitu assert dan require.
- Saat kita menggunakan assert akan memanggil `Fail()`, artinya eksekusi unit test akan tetap dilanjutkan.
- Sedangkan jika kita menggunakan `require`, jika pengecekan gagal, maka require akan memanggil `FailNow()`, artinya eksekusi unit test tidak akan dilanjutkan.

---

```go
package helper

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Yusril")
	assert.Equal(t, "Hello Yusril", result, "Result must be Hello Yusril")
	fmt.Println("Dieksekusi")
}
```

---

```go
package helper

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
  result := HelloWorld("Yusril")
  require.Equal(t, "Hell Yusril", result)
  fmt.Println("tidak dieksekusi")
}
```

## Skip Test

- Kadang dalam keadaan tertentu ingin membatalkan eksekusi unit test.
- Di Go-Lang juga kita bisa membatalkan eksekusi unit test jika kita mau.
- Untuk membatalkan unit test kita bisa menggunakan function `Skip()`.

### Kode : Skip Test

```go
func TestSkip(t *testing.T) {
  if runtime.GOOS == "linux" {
    t.Skip("Unit test tidak bisa jalan di Linux");
  }

  result := HelloWorld("Yusril")
  require.Equal(t, "Hello Yusril", result)
}
```

## Before dan After Test

- Biasanya dalam unit test, kadang kita ingin melakukan sesuatu sebelum dan setelah sebuah unit test dieksekusi.
- Jikalau kode yang kita lakukan sebelum dan setelah selalu sama antar unit test function, maka membuat manual di unit test functionnya adalah hal yang membosankan dan terlalu banyak kode duplikat jadiinya.
- Untungnya di Go-Lang terdaat fitur yang bernama `testing.M`.
- Fitur ini bernama Main, dimana digunakan untuk mengatur eksekusi unit test, namun hal ini juga bisa kita gunakan untuk melakukan Before dan After di unit test.

### `testing.M` Main Test

- Untuk mengatur eksekusi unit test, kita cukup membuat sebuah function benama `TestMain` dnegan parameter `testing.M`
- Jika terdapat function `TestMain` tersebut, maka secara otomatis Go-Lang akan mengeksekusi function ini tiap kali akan menjalankan unit test di sebuah package.
- Dengan ini kita bisa mengatur Before dan After unit test sesuai dengan yang kita mau.
- Ingat, function `TestMain` itu dieksekusi hanya sekali per Go-Lang package, bukan per tiap function unit test.

### Kode : `TestMain`

```go
package helper

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
  fmt.Println("Sebelum Unit Test")

  m.Run() // eksekusi semua unit test

  fmt.Println("Setelah Unit Test")
}
```

## Sub Test

- Go-Lang mendukung fitur pembuatan function unit test di dalam function unit test.
- Fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test dibahasa pemrograman yang lainnya.
- Untuk membuat sub test, kita bisa menggunakan function `Run()`.

### Kode : Membuat Sub Test

```go
func TestSubTest(t *testing.T) {
  t.Run("Yusril", func(t *testing.T) {
    result := HelloWorld("Yusril")
    require.Equal(t, "Hello Yusril", result)
  })

  t.Run("Arzaqi", func(t *testing.T) {
    result := HelloWorld("Arzaqi")
    require.Equal(t, "Hello Arzaqi", result)
  })
}
```

### Menjalankan Hanya Sub Test

- Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa gunakan perintah :
  - `go test -run TestNamaFunction`.
- Jika kita ingin menjalankan hanya salah satu sub test, kita bisa gunakan perintah :
  - `go test -run TestNamaFunction/NamaSubTest`.
- Atau untuk semua test semua test di semua function, kita bisa gunakan perintah :
  - `go test -run /NamaSubTest`

## Table Test

- Sebelumnya kita sudah belajar tentan sub test.
- Jika diperhatikan, sebenarnya dengan sub test, kita bisa membuat test secara dinamis.
- Dan fitur sub test ini, biasa digunakan oleh programmer Go-Lang untuk membuat test dengan konsep table test.
- Table test yaitu dimana kita menyediakan data berupa slice yang berisi parameter dan ekspetasi hasil dari unit test.
- Lalu slice tersebut kita iterasi menggunakan sub test.

### Kode : Table Test

```go
package helper

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorldTable(t *testing.T) {
	tests := [...]struct {
		name, request, expected string
	}{
		{
			name:     "HelloWold(Yusril)",
			request:  "Yusril",
			expected: "Hello Yusril",
		},
    {
      name: "HelloWorld(Bimo)",
      request: "Bimo",
      expected: "Hello Bimo",
    },
    {
      name: "HelloWorld(Arzaqi)",
      request: "Arzaqi",
      expected: "Hello Arzaqi",
    },
	}

  for _, test := range tests {
    t.Run(test.name, func(t *testing.T) {
      result := HelloWorld(test.request)
      require.Equal(t, test.expected, result)
    })
  }
}
```

## Mock

- Mock adalah object yang sudah kita program dengan ekspektasi tertentu sehingga ketika dipanggil, dia akan menghasilkan data yang sudah kita program diawal.
- Mock adalah salah sau teknik dalam unit testing, dimana kita bisa membuat mock object dari suatu object yang memang sulit untuk di testing.
- Misal kita ingin membuat unit test, namun ternyata ada kode program kita yang harus memanggil API Call ke thrid party service, dan belum tentu response nya sesuai dengan apa yang kita mau.
- PAda kasus seperti ini, cocok sekali untuk menggunakan mock object.

### Testify Mock

- Untuk membuat object, tidak ada fitur bawaan Go-Lang, namun kita bisa menggunakan library testify yang Sebelumnya kita gunakan untuk assertion.
- Testify mendukung pembuatan mock object, sehingga cocok untuk kita gunakan ketika ingin membuat mock object.
- Namun, perlu diperhatikan, jika desain kode program kita jelek, akan sulit untuk melakukan mocking, jadi pastikan kita melakukan pembuatan desain kode program kita dengan baik.
- Mari kita buat contoh kasus.

### Aplikasi Query Ke Database

- Kita akan coba contoh kasus dengan membuat contoh aplikasi golang yang melakukan query ke database.
- Dimana kita akan buat layer Service sebagai business logic, dan layer Repository sebagai jembatan ke database.
- Agar kode kita mudah untuk di test, disarankan agar membuat kontrak berupa `interface`.

### Kode : Category Entity

```go
package entity

type Category struct {
  Id, Name string
}
```

### Kode : Category Repository

```go
package repository

import "belajar-golang-unit-test/entity"

type CategoryRepository interface {
  FindId(id string) *entity.Category
}
```

### Kode : Category Service

```go
package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
)

type CategoryService struct {
  Repository repository.CategoryRepository;
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
  category := service.Repository.FindId(id)
  if category == nil {
    return category, errors.New("category Not Found")
  } else {
    return category, nil
  }
}
```

## Benchmark

- Selain unit test, Go-Lang testing package juga mendukung melakukan benchmark.
- Benchmark adalah mekanisme menghitung kecepatan performa kode aplikasi kita.
- Benchmark di Go-Lang dilakukan dengan cara secara otomatis melakukan iterasi kode yang kita panggil berkali-kali sampai waktu tertentu.
- Kita tidak perlu menentukan jumlah iterasi dan lamanya, karena itu sudah diatur oleh `testing.B` bawaan dari testing package.

### `testing.B` (Benchmark)

- `testing.B` adalah `struct` yang digunakan untuk melakukan benchmark.
- `testing.B` mirip dengan `testing.T`, terdapat function `Fail()`, `FailNow()`, `Error()`, `Fatal()` dan lain-lain.
- Yang membedakan, ada beberapa attribute dan function tambahan yang digunakan untuk melakukan benchmark.
- Salah satunya adalah attrubute `N`, ini digunakna untuk melakukan total iterasi sebuah benchmark.

### Cara Kerja Benchmark

- Cara kerja benchmark di Go-Lang sangat sederhana.
- Gimana kita hanya perlu membuat perulangan sejumlah `N` attribute.
- Nanti secara otomatis Go-Lang akan melakukan eksekusi sejumlah perulangan yang ditentukkan secara otomatis, lalu mendeteksi berapa lama proses tersebut berjalan, dan disimpulkan performa benchmark nya dalam waktu.

## Benchmark Function

- Mirip seperti unit test, untuk benchmark pun, di Go-Lang sudah ditentukan nama functionnya, harus diawali dengan kata `Benchmark`, misal `BenchmarkHelloWorld`, `BenchmarXxx`.
- Selain itu, harus memiliki parameter `b *testing.B`.
- Dan tidak boleh mengembalikan return value.
- Untuk nama file benchmark, sama seperti unit test, diakhiri dengan `_test`, misal `hello_world_test.go`.

### Kode : Membuat Benchmark Function

```go
func BenchmarkHelloWorld(b *testing.B) {
  for i := 0; i < b.N; i++ {
    HelloWorld("Yusril")
  }
}
```

### Menjalankan Benchmark

- Untuk menjalankan seluruh benchmark di module, kita bisa menggunakan perintah sama seperti test, namun ditambahkan parameter bench :
  - `go test -v -bench=`.
- Jika kita hanya ingin menjalankan benchmark tanpa unit test, kita bisa gunakan perintah :
  - `go test -v -run=NotMathUnitTest -bench=`.
- Kode diatas selain menjalankan benchmark tertentu, kita bisa gunakan perintah :
  - `go test -v -run=NotMatchUnitTest -bench=BenchmarkTest`.
- Jika kita menjalankan benchmark di root module dan ingin semua module dijalankan, kita bisa gunakan perintah :
  - `go test -v -bench=./...`

## Sub Benchmark

- Sama seperti `testing.T`, di `testing.B` juga kita bisa membuat sub benchmark menggunakan function `Run()`.

### Kode : Membuat Sub Benchmark

```go
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
```

### Menjalankan Hanya Sub Benchmark

- Saat kita menjalankan benchmark function, maka semua sub benchmark akan bejalan.
- Namun jika kita ingin menjalankan salah satu sub benchmark saja, kita bisa gunakan perintah :
  - `go test -v -bench=BenchmarkNama/NamaSub`

### Table Benchmark

- Sama seperti di unit test, programmer Go-Lang terbiasa membuat table benchmark juga.
- Ini digunakan agar kita bisa mudah melakukan performance test dengan kombinasi data berbeda-beda tanpa harus membuat benyak benchmark function.

### Kode : Table Benchmark

```go
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name, request string
	}{
		{
			name:    "HelloWorld(Yusril)",
			request: "Yusril",
		},
		{
			name:    "HelloWorld(Arzaqi)",
			request: "Arzaqi",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
```
