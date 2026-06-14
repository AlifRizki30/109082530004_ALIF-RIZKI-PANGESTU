# <h1 align="center">sLaporan Praktikum Modul 14 - ... </h1>
<p align="center">[ALIF RIZKI PANGESTU] - [109082530004]</p>

## Unguided 

### 1A. [Soal]
#### soal1.go

```go
package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIdx := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		rumah := make([]int, m)

		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		selectionSort(rumah)

		for j := 0; j < m; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[j])
		}
		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%2014/soal%201a.go)Program ini digunakan untuk mengurutkan nomor rumah kerabat Hercules di setiap daerah menggunakan metode Selection Sort. Pertama, program meminta jumlah daerah yang akan diproses (n). Untuk setiap daerah, program membaca jumlah rumah kerabat (m) dan nomor-nomor rumahnya, lalu menyimpannya ke dalam array. Setelah itu, nomor rumah diurutkan dari yang paling kecil ke yang paling besar dengan cara mencari angka terkecil yang belum terurut dan menukarnya ke posisi yang tepat. Setelah proses pengurutan selesai, program menampilkan nomor rumah yang sudah terurut untuk setiap daerah dalam satu baris. Dengan demikian, output yang dihasilkan adalah daftar nomor rumah yang tersusun rapi secara menaik di masing-masing daerah.


## Unguided 

### 1B. [Soal]
#### soal2.go

```go
package main

import "fmt"

func selectionSortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func selectionSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		max := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		arr[i], arr[max] = arr[max], arr[i]
	}
}

func main() {
	var n int

	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int

		fmt.Printf("\nDaerah ke-%d\n", i+1)
		fmt.Print("Masukkan jumlah rumah kerabat: ")
		fmt.Scan(&m)

		ganjil := []int{}
		genap := []int{}

		fmt.Println("Masukkan nomor rumah kerabat:")
		for j := 0; j < m; j++ {
			var x int
			fmt.Scan(&x)

			if x%2 == 1 {
				ganjil = append(ganjil, x)
			} else {
				genap = append(genap, x)
			}
		}

		selectionSortAsc(ganjil)
		selectionSortDesc(genap)

		fmt.Print("Urutan rumah kerabat: ")

		pertama := true
		for _, x := range ganjil {
			if !pertama {
				fmt.Print(" ")
			}
			fmt.Print(x)
			pertama = false
		}

		for _, x := range genap {
			if !pertama {
				fmt.Print(" ")
			}
			fmt.Print(x)
			pertama = false
		}

		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%2014/soal%201b.go)
Program ini digunakan untuk mengurutkan nomor rumah kerabat Hercules agar ia tidak perlu terlalu sering menyeberang jalan. Program terlebih dahulu meminta jumlah daerah dan jumlah rumah kerabat di setiap daerah. Setiap nomor rumah yang dimasukkan akan dipisahkan menjadi dua kelompok, yaitu nomor ganjil dan nomor genap. Nomor rumah ganjil kemudian diurutkan dari yang paling kecil ke yang paling besar menggunakan algoritma Selection Sort, sedangkan nomor rumah genap diurutkan dari yang paling besar ke yang paling kecil. Setelah proses pengurutan selesai, program menampilkan semua nomor rumah ganjil terlebih dahulu, kemudian diikuti nomor rumah genap yang sudah terurut. Dengan cara ini, nomor rumah pada satu sisi jalan (ganjil) dikunjungi lebih dulu secara berurutan, lalu dilanjutkan ke sisi lainnya (genap), sehingga jumlah penyeberangan jalan dapat diminimalkan.


## Unguided 

### 2A. [Soal]
#### soal3.go

```go
package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func main() {
	var x int
	var data []int

	fmt.Println("Masukkan bilangan (akhiri dengan bilangan negatif):")

	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		data = append(data, x)
	}

	insertionSort(data)

	for i := 0; i < len(data); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(data[i])
	}
	fmt.Println()

	if len(data) <= 2 {
		fmt.Println("Data berjarak tetap")
		return
	}

	jarak := data[1] - data[0]
	tetap := true

	for i := 2; i < len(data); i++ {
		if data[i]-data[i-1] != jarak {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%2014/soal%202a.go)
Program ini digunakan untuk membaca sekumpulan bilangan bulat yang dimasukkan pengguna hingga ditemukan bilangan negatif sebagai tanda berhenti. Hanya bilangan yang bernilai nol atau positif yang disimpan ke dalam array. Setelah semua data masuk, program mengurutkan bilangan tersebut dari yang terkecil ke yang terbesar menggunakan algoritma Insertion Sort, yaitu dengan menyisipkan setiap elemen ke posisi yang tepat pada bagian array yang sudah terurut. Selanjutnya, program menampilkan hasil pengurutan dan memeriksa apakah selisih antara setiap dua bilangan yang berurutan selalu sama. Jika semua selisihnya sama, program menampilkan pesan Data berjarak x sesuai nilai jaraknya, sedangkan jika ada selisih yang berbeda maka program menampilkan Data berjarak tidak tetap. Dengan demikian, program tidak hanya mengurutkan data tetapi juga mengecek apakah data tersebut membentuk pola barisan dengan jarak yang konstan.



## Unguided 

### 2B. [Soal]
#### soal3.go

```go
package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax + 1]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i <= n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n > 0 {
		maxRating := pustaka[1].rating
		for i := 2; i <= n; i++ {
			if pustaka[i].rating > maxRating {
				maxRating = pustaka[i].rating
			}
		}

		for i := 1; i <= n; i++ {
			if pustaka[i].rating == maxRating {
				fmt.Printf("%s, %s, %s, %d\n", pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit, pustaka[i].tahun)
			}
		}
	}
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 2; i <= n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 1 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 1; i <= limit; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri := 1
	kanan := n
	found := -1

	for kiri <= kanan && found == -1 {
		med := (kiri + kanan) / 2
		if pustaka[med].rating == r {
			found = med
		} else if pustaka[med].rating < r {
			kanan = med - 1
		} else {
			kiri = med + 1
		}
	}

	if found != -1 {
		fmt.Printf("%s, %s, %s, %d, %d, %d\n", pustaka[found].judul, pustaka[found].penulis, pustaka[found].penerbit, pustaka[found].tahun, pustaka[found].eksemplar, pustaka[found].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var n, targetRating int
	var pustaka DaftarBuku

	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)
	fmt.Scan(&targetRating)

	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	CariBuku(pustaka, n, targetRating)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%2014/soal%202b.go)
Program ini digunakan untuk mengelola data buku pada sebuah perpustakaan. Pertama, program membaca sejumlah data buku yang berisi ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating. Setelah semua data dimasukkan, program mencari dan menampilkan buku dengan rating tertinggi sebagai buku terfavorit. Selanjutnya, data buku diurutkan berdasarkan rating dari yang terbesar ke yang terkecil menggunakan algoritma Insertion Sort, kemudian ditampilkan maksimal lima judul buku dengan rating tertinggi. Terakhir, program melakukan pencarian buku berdasarkan rating yang dimasukkan pengguna menggunakan metode Binary Search pada data yang sudah terurut. Jika ditemukan, program akan menampilkan informasi lengkap buku tersebut, sedangkan jika tidak ditemukan akan muncul pesan bahwa tidak ada buku dengan rating yang dicari. Dengan demikian, program dapat digunakan untuk menyimpan, mengurutkan, menampilkan buku favorit, dan mencari buku berdasarkan ratingnya.

