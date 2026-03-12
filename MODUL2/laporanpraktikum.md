# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[ALIF RIZKI PANGESTU] - [109082530004]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"
func main() {
	var (
		satu, dua, tiga string
		temp string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL2/OUTPUT/OUTPUTSOAL1.png)
codingan tersebut tersebut berfungsi untuk menerima tiga buah input string dari pengguna yang disimpan dalam variabel satu, dua, dan tiga menggunakan fmt.Scanln. Setelah itu program menampilkan urutan nilai awal dari ketiga variabel tersebut menggunakan fmt.Println. Selanjutnya dilakukan proses pertukaran atau pergeseran nilai dengan bantuan variabel sementara temp, di mana nilai satu dipindahkan ke temp, nilai dua dipindahkan ke satu, nilai tiga dipindahkan ke dua, dan nilai yang ada di temp dipindahkan ke tiga. Setelah proses pertukaran selesai, program menampilkan kembali hasil akhir dari ketiga variabel sehingga urutan nilainya berubah dari urutan awal.

## Unguided 

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"
func main() {
	var w1, w2, w3, w4 string
	var berhasil bool
	fmt.Println("urutkan warna sesuai rules")
	berhasil = true
	i:=1
	for  i <= 5 {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&w1, &w2, &w3, &w4)

		if w1 != "merah" || w2 != "kuning" || w3 != "hijau" || w4 != "ungu" {
			berhasil = false
		}
		i++
	}

	fmt.Println("BERHASIL:", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL2/OUTPUT/OUTPUTSOAL2.png)
codingan tersebut digunakan untuk mengecek apakah urutan warna yang dimasukkan pengguna sudah sesuai dengan aturan yang ditentukan, var merah, kuning, hijau, ungu. Pada awal program dideklarasikan empat var string w1, w2, w3, w4 untuk menyimpan input warna serta satu variabel boolean berhasil yang bernilai awal true. Program kemudian menampilkan pesan kepada pengguna untuk memasukkan urutan warna dan melakukan perulangan for sebanyak lima kali sebagai percobaan input. Pada setiap percobaan, pengguna diminta memasukkan empat warna yang kemudian diperiksa menggunakan kondisi if. Jika salah satu warna tidak sesuai dengan urutan yang ditentukan (merah, kuning, hijau, ungu), maka variabel berhasil akan berubah menjadi false. Setelah seluruh percobaan selesai, program akan menampilkan hasil akhir berupa nilai berhasil yang menunjukkan apakah semua input yang dimasukkan sesuai dengan aturan atau tidak.


## Unguided 

### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"
func main() {
	var gram int
	var Kg, Sisa int
	var ongkosKg, ongkosSisa, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gram)

	Kg = gram / 1000
	Sisa = gram % 1000
	ongkosKg = Kg * 10000

	if Kg > 10 {
		ongkosSisa = 0
	} else {
		if Sisa >= 500 {
			ongkosSisa = Sisa * 5
		} else {
			ongkosSisa = Sisa * 15
		}
	}
	total = ongkosKg + ongkosSisa

	fmt.Println("Detail berat:", Kg, "Kg+", Sisa, "gr")
	fmt.Println("Detail ongkos: Rp. ", ongkosKg, "+ Rp. ", ongkosSisa)
	fmt.Println("Total ongkos Rp.", total)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL2/OUTPUT/OUTPUTSOAL3.png)
[penjelasan]







