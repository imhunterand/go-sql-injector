# Go SQL Injector

Go SQL Injector adalah alat keamanan berbasis Go yang dirancang untuk mendeteksi dan mengeksploitasi celah keamanan SQL Injection pada aplikasi web. Alat ini memungkinkan pengguna untuk mengekstrak nama tabel dan kolom dari basis data yang rentan, membantu dalam pengujian penetrasi dan audit keamanan.

## Fitur

- **Deteksi Tabel**: Mendeteksi nama tabel yang ada dalam basis data target melalui celah SQL Injection.
- **Deteksi Kolom**: Mendeteksi nama kolom dalam tabel yang terdeteksi melalui celah SQL Injection.
- **Antarmuka Pengguna Berbasis CLI**: Menyediakan antarmuka command-line yang mudah digunakan untuk melakukan injeksi dan ekstraksi data.
- **Modular dan Mudah Dikembangkan**: Struktur kode yang modular memudahkan pengembangan dan penambahan fitur baru.

## Struktur Proyek
```
go-sql-injector/
├── main.go
├── injector.go
├── extractor.go
├── utils.go
└── go.mod
```
## Instalasi

1. Clone repository ini:
```
git clone https://github.com/username/go-sql-injector.git
```
2. Navigasi ke direktori proyek:
```
cd go-sql-injector
```
3. Install dependensi:
```
go mod tidy
```
## Penggunaan
1. Jalankan `main.go` dengan URL target:
```
go run main.go <url>
```
Gantilah `<url>` dengan URL target yang rentan terhadap SQL Injection.

2. Ikuti petunjuk pada antarmuka CLI untuk mendeteksi tabel dan kolom dalam basis data target.

## Contoh

Setelah mengikuti petunjuk penggunaan, alat ini akan mengekstrak nama tabel dan kolom dari basis data target yang rentan terhadap SQL Injection.

## Kontribusi

Kami menyambut kontribusi dari pengembang lain. Silakan fork repository ini dan ajukan pull request dengan penjelasan tentang perubahan yang Anda buat.
