# <h1 align="center">sLaporan Praktikum Modul 10 - ... </h1>
<p align="center">[ALIF RIZKI PANGESTU] - [109082530004]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var n int
	var berat [1000]float64
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}

		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", min)
	fmt.Printf("Berat terbesar: %.2f\n", max)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%2010/OUTPUT/soal1.png)Program Go pada gambar digunakan untuk mencari berat kelinci terkecil dan terbesar dari sejumlah data yang dimasukkan pengguna. Program dimulai dengan meminta input jumlah anak kelinci, kemudian setiap berat kelinci disimpan ke dalam array bertipe float64. Setelah semua data dimasukkan, program menggunakan perulangan dan percabangan if untuk membandingkan setiap nilai sehingga dapat menentukan nilai minimum (min) dan maksimum (max). Hasil akhirnya ditampilkan dalam bentuk berat terkecil dan berat terbesar. Program ini menerapkan konsep dasar pemrograman seperti input/output, array, perulangan, dan percabangan dalam bahasa Go (Golang).


## Unguided 

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func main() {
	var x, y int
	var ikan [1000]float64

	fmt.Scan(&x, &y)
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}
	index := 0

	for index < x {
		total := 0.0
		jumlahIkan := 0

		for j := 0; j < y && index < x; j++ {
			total += ikan[index]
			jumlahIkan++
			index++
		}

		fmt.Printf("%.2f ", total)

	}

	fmt.Println()
	index = 0

	for index < x {
		total := 0.0
		jumlahIkan := 0

		for j := 0; j < y && index < x; j++ {
			total += ikan[index]
			jumlahIkan++
			index++
		}

		rata := total / float64(jumlahIkan)
		fmt.Printf("%.2f ", rata)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%2010/OUTPUT/soal2.png)Program Go pada gambar digunakan untuk mengelompokkan data berat ikan dan menghitung rata-rata berat pada setiap kelompok. Program dimulai dengan menerima input jumlah data ikan (x) dan jumlah anggota tiap kelompok (y), kemudian seluruh data berat ikan disimpan ke dalam array bertipe float64. Setelah itu, program menggunakan perulangan untuk menjumlahkan beberapa data ikan sesuai ukuran kelompok, lalu menghitung rata-ratanya dengan membagi total berat dengan jumlah ikan pada kelompok tersebut. Hasil rata-rata setiap kelompok kemudian ditampilkan ke layar. Program ini menerapkan konsep array, input/output, perulangan bertingkat, serta pengolahan data numerik menggunakan bahasa Go (Golang).


## Unguided 

### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}

		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}
func rerata(arrBerat arrBalita, n int) float64 {
	total := 0.0

	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int
	var min, max, rata float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d : ", i+1)
		fmt.Scan(&data[i])
	}
	hitungMinMax(data, n, &min, &max)
	rata = rerata(data, n)

	// Output
	fmt.Printf("\nBerat balita minimum : %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum : %.2f kg\n", max)
	fmt.Printf("\nRerata berat balita : %.2f kg\n", rata)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%2010/OUTPUT/soal3.png)Program Go pada gambar digunakan untuk mengolah data berat balita dengan mencari berat minimum, berat maksimum, dan rata-rata berat balita. Program menggunakan array untuk menyimpan data berat yang dimasukkan pengguna, kemudian memanfaatkan fungsi hitungMinMax() untuk menentukan nilai terkecil dan terbesar melalui proses perbandingan setiap data dalam array. Selain itu, fungsi rerata() digunakan untuk menghitung rata-rata berat balita dengan menjumlahkan seluruh data lalu membaginya dengan jumlah balita. Setelah proses selesai, program menampilkan hasil berat minimum, maksimum, dan rata-rata dalam satuan kilogram. Program ini menerapkan konsep fungsi, array, perulangan, percabangan, pointer, serta pengolahan data numerik dalam bahasa Go (Golang).

