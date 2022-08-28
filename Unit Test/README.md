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

##
