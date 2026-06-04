# <h1 align="center">sLaporan Praktikum Modul 10 - ... </h1>
<p align="center">[ALIF RIZKI PANGESTU] - [109082530004]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var suara int
	totalMasuk := 0
	suaraSah := 0
	var calon [21]int

	for {
		fmt.Scan(&suara)

		totalMasuk++

		if suara == 0 {
			break
		}

		if suara >= 1 && suara <= 20 {
			suaraSah++
			calon[suara]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	for i := 1; i <= 20; i++ {
		if calon[i] > 0 {
			fmt.Printf("%d: %d\n", i, calon[i])
		}
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%2012/soal1.go)
Program Go tersebut digunakan untuk menghitung hasil pemilihan ketua RT berdasarkan data suara yang dimasukkan pengguna. Program membaca setiap suara satu per satu hingga ditemukan angka 0 sebagai penanda akhir input. Setiap data yang masuk dihitung sebagai suara masuk, kemudian divalidasi apakah berada pada rentang 1 sampai 20. Jika valid, suara tersebut dihitung sebagai suara sah dan jumlah suara untuk calon yang bersangkutan ditambahkan pada array calon. Setelah seluruh data selesai dibaca, program menampilkan jumlah total suara yang masuk, jumlah suara yang sah, serta daftar calon yang memperoleh suara beserta jumlah suara yang didapatkan. Dengan menggunakan array sebagai penyimpan frekuensi suara setiap calon, program dapat melakukan perhitungan suara secara sederhana dan efisien. 

## Unguided 

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func main() {
	var x int
	totalMasuk := 0
	suaraSah := 0

	var suara [21]int

	for {
		fmt.Scan(&x)

		totalMasuk++

		if x == 0 {
			break
		}

		if x >= 1 && x <= 20 {
			suaraSah++
			suara[x]++
		}
	}

	ketua := 1
	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			ketua = i
		}
	}

	wakil := -1
	for i := 1; i <= 20; i++ {
		if i == ketua {
			continue
		}

		if wakil == -1 ||
			suara[i] > suara[wakil] ||
			(suara[i] == suara[wakil] && i < wakil) {
			wakil = i
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%2012/soal2.go)
Program Go tersebut digunakan untuk menentukan pasangan Ketua RT dan Wakil Ketua RT berdasarkan hasil pemungutan suara yang dimasukkan pengguna. Program membaca data suara satu per satu hingga ditemukan angka 0 sebagai penanda akhir input, kemudian menghitung jumlah suara masuk dan suara sah, yaitu suara yang berada pada rentang nomor calon 1 sampai 20. Setiap suara sah disimpan dan dihitung menggunakan array sehingga jumlah suara yang diperoleh masing-masing calon dapat diketahui. Setelah seluruh data diproses, program mencari calon dengan jumlah suara terbanyak sebagai Ketua RT, lalu mencari calon dengan jumlah suara terbanyak berikutnya sebagai Wakil Ketua RT. Jika terdapat beberapa calon dengan jumlah suara yang sama, program secara otomatis memilih calon dengan nomor yang lebih kecil sebagai pemenang. Terakhir, program menampilkan jumlah suara masuk, jumlah suara sah, serta nomor calon yang terpilih sebagai Ketua RT dan Wakil Ketua RT. 


## Unguided 

### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	for i := 0; i < n; i++ {
		if data[i] == k {
			return i
		}
	}
	return -1
}

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	isiArray(n)

	idx := posisi(n, k)

	if idx == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(idx)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%2012/soal3.go)
Program Go tersebut digunakan untuk mencari apakah suatu bilangan k terdapat di dalam kumpulan data integer yang telah terurut membesar. Program diawali dengan membaca jumlah data n dan nilai yang akan dicari k, kemudian memanggil prosedur isiArray() untuk mengisi array global data sebanyak n elemen. Setelah data tersimpan, fungsi posisi() melakukan pencarian secara berurutan (linear search) dari indeks pertama hingga terakhir untuk menemukan nilai k. Jika nilai tersebut ditemukan, fungsi akan mengembalikan posisi atau indeks kemunculannya yang dihitung mulai dari 0, sedangkan jika tidak ditemukan fungsi mengembalikan nilai -1. Pada bagian utama program, hasil pencarian disimpan dalam variabel idx; apabila nilainya -1, program menampilkan pesan "TIDAK ADA", dan jika ditemukan maka program menampilkan indeks posisi bilangan tersebut dalam array. Program ini menunjukkan penggunaan array, prosedur, fungsi, perulangan, dan percabangan untuk menyelesaikan masalah pencarian data secara sederhana. 

