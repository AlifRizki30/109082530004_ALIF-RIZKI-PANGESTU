# <h1 align="center">sLaporan Praktikum Modul 9 - ... </h1>
<p align="center">[ALIF RIZKI PANGESTU] - [109082530004]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y float64
}

type Lingkaran struct {
	pusat Titik
	r     float64
}

func jarak(a, b Titik) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func diDalam(c Lingkaran, t Titik) bool {
	return jarak(c.pusat, t) <= c.r
}

func main() {
	var c1, c2 Lingkaran
	var t Titik

	fmt.Scan(&c1.pusat.x, &c1.pusat.y, &c1.r)
	fmt.Scan(&c2.pusat.x, &c2.pusat.y, &c2.r)
	fmt.Scan(&t.x, &t.y)

	in1 := diDalam(c1, t)
	in2 := diDalam(c2, t)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1]https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%209/OUTPUT/SOAL1.png
Kode Go tersebut intinya digunakan untuk mengecek apakah sebuah titik berada di dalam suatu lingkaran atau tidak. Pertama, dibuat tipe data Titik untuk menyimpan koordinat x dan y, lalu tipe Lingkaran yang punya pusat dan jari-jari. Setelah itu ada fungsi jarak yang menghitung jarak antara dua titik menggunakan rumus matematika. Fungsi berikutnya, diDalam, dipakai untuk menentukan apakah suatu titik berada di dalam lingkaran dengan cara membandingkan jarak titik ke pusat lingkaran dengan jari-jarinya—kalau jaraknya lebih kecil atau sama, berarti titik itu ada di dalam. Dari output terminalnya, program ini menerima beberapa input titik dan lingkaran, lalu menampilkan apakah titik tersebut berada di dalam lingkaran 1, lingkaran 2, atau bahkan keduanya.

## Unguided 

### 2. [Soal]
#### soal2.go

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Semua:", arr)

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Scan(&x)
	fmt.Print("Kelipatan indeks ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
	var idx int
	fmt.Scan(&idx)
	arr = append(arr[:idx], arr[idx+1:]...)
	fmt.Println("Setelah hapus:", arr)

	sum := 0
	for _, v := range arr {
		sum += v
	}
	mean := float64(sum) / float64(len(arr))
	fmt.Println("Rata-rata:", mean)
	var total float64
	for _, v := range arr {
		total += math.Pow(float64(v)-mean, 2)
	}
	std := math.Sqrt(total / float64(len(arr)))
	fmt.Println("Std Dev:", std)

	var cari int
	fmt.Scan(&cari)
	count := 0
	for _, v := range arr {
		if v == cari {
			count++
		}
	}
	fmt.Println("Frekuensi:", count)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1]https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%209/OUTPUT/SOAL2.png
Kode Go ini digunakan untuk mengolah sekumpulan angka yang dimasukkan oleh pengguna. Awalnya program meminta jumlah data (n), lalu membaca satu per satu angka dan menyimpannya ke dalam array. Setelah itu, program menampilkan semua isi array, kemudian mengambil dan menampilkan elemen pada indeks ganjil (posisi ke-2, ke-4, dst) dan indeks genap (posisi ke-1, ke-3, dst). Selanjutnya, program juga menampilkan elemen yang berada pada kelipatan indeks tertentu (misalnya indeks ke-2), lalu melakukan penghapusan pada salah satu elemen dari array, dan akhirnya menghitung serta menampilkan nilai rata-rata dari data yang tersisa. Intinya, program ini memperlihatkan berbagai operasi dasar pada array seperti input data, filtering berdasarkan indeks, manipulasi (hapus data), dan perhitungan sederhana seperti rata-rata.

## Unguided 

### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	pemenang := []string{}
	i := 1

	for {
		var a, b int
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&a, &b)

		if a < 0 || b < 0 {
			break
		}

		if a > b {
			pemenang = append(pemenang, klubA)
		} else if b > a {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
		i++
	}

	for i, v := range pemenang {
		fmt.Printf("Hasil %d : %s\n", i+1, v)
	}

	fmt.Println("Pertandingan selesai")
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1]https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%209/OUTPUT/SOAL3.png
Kode Go ini digunakan untuk mensimulasikan pertandingan antara dua klub. Awalnya, pengguna diminta memasukkan nama klub A dan klub B. Setelah itu, program akan terus meminta input skor untuk setiap pertandingan (misalnya skor klub A dan klub B) secara berulang, sambil menghitung nomor pertandingan. Jika salah satu skor yang dimasukkan bernilai negatif, perulangan akan berhenti. Selama proses tersebut, program menentukan siapa pemenang tiap pertandingan—kalau skor A lebih besar maka klub A menang, kalau B lebih besar maka klub B menang, dan kalau sama berarti seri (draw). Semua hasil pertandingan kemudian disimpan dalam sebuah array, dan di akhir program ditampilkan satu per satu hasilnya, jadi kita bisa melihat riwayat siapa yang menang di setiap pertandingan dengan cara yang sederhana.

## Unguided 

### 4. [Soal]
#### soal4.go

```go
package main

import "fmt"

const NMAX = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	*n = 0
	for {
		var ch rune
		fmt.Scanf("%c\n", &ch)

		if ch == '.' || *n >= NMAX {
			break
		}
		t[*n] = ch
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var t tabel
	var n int

	fmt.Println("Masukkan teks (akhiri dengan .):")
	isiArray(&t, &n)

	fmt.Print("Teks: ")
	cetakArray(t, n)

	if palindrom(t, n) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}

	balikanArray(&t, n)

	fmt.Print("Reverse teks: ")
	cetakArray(t, n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1]https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%209/OUTPUT/SOAL4.png
Kode Go ini digunakan untuk membaca teks dari pengguna satu per satu huruf, lalu mengolahnya seperti array karakter. Program akan terus menerima input huruf sampai pengguna memasukkan tanda titik (.) sebagai tanda berhenti atau jika kapasitas sudah penuh. Semua huruf yang dimasukkan disimpan ke dalam array, lalu ditampilkan kembali sebagai teks utuh. Setelah itu, program mengecek apakah teks tersebut merupakan palindrom (dibaca sama dari depan dan belakang), dan juga menampilkan versi kebalikannya (reverse). Jadi secara sederhana, program ini berfungsi untuk input teks manual per karakter, lalu mengecek apakah kata tersebut simetris dan menampilkan hasil aslinya serta hasil yang sudah dibalik.