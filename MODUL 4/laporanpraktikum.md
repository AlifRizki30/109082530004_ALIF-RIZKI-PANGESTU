# <h1 align="center">sLaporan Praktikum Modul 4 - ... </h1>
<p align="center">[ALIF RIZKI PANGESTU] - [109082530004]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main
import (
	"fmt"
)

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if a < c || b < d {
		fmt.Println("Input tidak valid (harus a ≥ c dan b ≥ d)")
		return
	}
	p1 := permutation(a, c)
	c1 := combination(a, c)

	p2 := permutation(b, d)
	c2 := combination(b, d)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%204/OUTPUT/soal%201.png)
Kode Go tersebut digunakan untuk menghitung permutasi dan kombinasi dari dua pasang bilangan input. Program mendefinisikan fungsi faktorial(n int) untuk menghitung nilai faktorial (meskipun bagian perhitungannya tidak terlihat lengkap di potongan ini), lalu menggunakan fungsi permutation(n, r) dengan rumus ( n! / (n-r)! ) dan combination(n, r) dengan rumus ( n! / (r!(n-r)!) ). Pada fungsi main, program membaca empat input bilangan (a, b, c, d), kemudian melakukan validasi agar nilai a ≥ c dan b ≥ d; jika tidak, program menampilkan pesan error. Jika valid, program menghitung nilai permutasi dan kombinasi untuk pasangan `(a, c) dan (b, d), lalu menampilkan hasilnya ke terminal.

## Unguided 

### 2. [Soal]
#### soal2.go

```go
package main
import (
	"fmt"
)

func hitungSkor(waktu []int, soal *int, skor *int) {
	*soal = 0
	*skor = 0
	for _, w := range waktu {
		if w <= 301 {
			*soal++
			*skor += w
		}
	}
}
func main() {
	var nama string
	var pemenang string
	maxSoal := -1
	minWaktu := 1<<31 - 1

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}
		waktu := make([]int, 8)
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		var soal, skor int
		hitungSkor(waktu, &soal, &skor)
		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			maxSoal = soal
			minWaktu = skor
			pemenang = nama
		}
	}
	fmt.Println(pemenang, maxSoal, minWaktu)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1]https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%204/OUTPUT/soal%202.png
Program Go tersebut digunakan untuk menentukan pemenang berdasarkan jumlah soal yang diselesaikan dan total waktu tercepat. Di dalam main, program membaca input berupa nama peserta secara berulang hingga pengguna mengetik Selesai. Untuk setiap peserta, program mengambil 8 data waktu pengerjaan soal yang disimpan dalam slice waktu, kemudian memanggil fungsi hitungSkor untuk menghitung jumlah soal yang berhasil diselesaikan (soal) dan total waktu (skor). Program menyimpan nilai maksimum soal (maxSoal) dan waktu minimum (minWaktu) sebagai acuan, lalu membandingkan setiap peserta: jika jumlah soal lebih banyak atau sama tetapi dengan waktu lebih cepat, maka peserta tersebut menjadi pemenang. Setelah semua input selesai, program menampilkan nama pemenang beserta jumlah soal dan total waktunya.


## Unguided 

### 3. [Soal]
#### soal3.go

```go
package main

import (
	"fmt"
)

func cetakDeret(n int) {
	for {
		fmt.Print(n, " ")

		if n == 1 {
			break
		}

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1]https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%204/OUTPUT/soal%203.png
Program Go tersebut digunakan untuk mencetak deret bilangan berdasarkan aturan Collatz (3n + 1). Fungsi cetakDeret(n int) akan menampilkan nilai n secara berulang menggunakan loop tak hingga (for), lalu berhenti ketika nilai n mencapai 1. Di setiap iterasi, program mengecek apakah n genap atau ganjil: jika genap maka n dibagi 2 (n = n / 2), sedangkan jika ganjil maka dihitung dengan rumus n = 3*n + 1. Fungsi main digunakan untuk membaca input bilangan dari pengguna, kemudian memanggil fungsi cetakDeret untuk menampilkan seluruh urutan angka hingga berakhir di 1, seperti yang terlihat pada output deret yang terus berubah sesuai aturan tersebut.










