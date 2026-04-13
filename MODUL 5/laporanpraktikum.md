# <h1 align="center">sLaporan Praktikum Modul 5 - ... </h1>
<p align="center">[ALIF RIZKI PANGESTU] - [109082530004]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func main() {
	n := 10
	fmt.Println("Deret Fibonacci sampai suku ke-", n)
	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%205/OUTPUT/SOAL1.png)
Program Go tersebut digunakan untuk menampilkan deret Fibonacci hingga suku ke-n dengan memanfaatkan konsep rekursif. Program mendefinisikan sebuah fungsi fibonacci(n int) yang menghitung nilai suku ke-n, di mana jika n bernilai 0 maka mengembalikan 0, dan jika n bernilai 1 maka mengembalikan 1, sedangkan untuk nilai lainnya dihitung dari penjumlahan dua suku sebelumnya yaitu fibonacci(n-1) + fibonacci(n-2). Pada fungsi utama (main), program menentukan atau menerima input nilai n, kemudian menggunakan perulangan untuk memanggil fungsi fibonacci dari indeks 0 hingga n, sehingga menghasilkan dan menampilkan deret Fibonacci secara berurutan. Program ini menunjukkan bagaimana rekursi bekerja dengan memecah masalah menjadi sub-masalah yang lebih kecil hingga mencapai kondisi dasar (base case).


## Unguided 

### 2. [Soal]
#### soal2.go

```go
package main
import "fmt"
func printPattern(n int, current int) {
	if current > n {
		return
	}
	for i := 0; i < current; i++ {
		fmt.Print("*")
	}
	fmt.Println()

	printPattern(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	printPattern(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%205/OUTPUT/SOAL2.png)
Program Go pada soal nomor 2 bertujuan untuk menampilkan pola bintang berbentuk segitiga menggunakan konsep rekursif. Program menerima sebuah input bilangan N dari pengguna yang menentukan jumlah baris pola yang akan ditampilkan. Selanjutnya, terdapat fungsi rekursif yang bekerja dengan cara mencetak sejumlah bintang sesuai nilai baris saat ini, dimulai dari 1 hingga N, kemudian memanggil dirinya sendiri dengan nilai yang meningkat hingga mencapai batas. Setiap pemanggilan fungsi akan mencetak satu baris bintang sebelum melanjutkan ke baris berikutnya, sehingga terbentuk pola bertingkat dari satu bintang hingga N bintang. Konsep ini memanfaatkan rekursi sebagai alternatif dari perulangan untuk menghasilkan pola berulang, di mana setiap langkah bergantung pada langkah sebelumnya hingga mencapai kondisi berhenti.



## Unguided 

### 3. [Soal]
#### soal3.go

```go
package main
import "fmt"
func faktor(n int, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Print(i, " ")
	}

	faktor(n, i+1)
}
func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	fmt.Print("Faktor dari ", n, ": ")
	faktor(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%205/OUTPUT/SOAL3.png)
Program Go pada soal nomor 3 bertujuan untuk menampilkan faktor dari suatu bilangan dengan menggunakan konsep rekursif. Program menerima input berupa sebuah bilangan bulat positif N, kemudian menggunakan fungsi rekursif untuk mengecek setiap bilangan dari 1 hingga N apakah habis membagi N atau tidak. Jika suatu bilangan memenuhi kondisi N % i == 0, maka bilangan tersebut dicetak sebagai faktor. Proses ini dilakukan berulang dengan cara fungsi memanggil dirinya sendiri dengan nilai yang terus bertambah hingga mencapai batas N, lalu berhenti ketika kondisi akhir terpenuhi. Program ini menunjukkan penerapan rekursi sebagai teknik pengulangan, di mana fungsi terus memanggil dirinya sendiri sampai mencapai kondisi berhenti (base case) sehingga seluruh faktor dari bilangan tersebut dapat ditampilkan secara berurutan


## Unguided 

### 4. [Soal]
#### soal4.go

```go
package main

import "fmt"

func barisan(n int) {
	if n == 0 {
		return
	}

	fmt.Print(n, " ")
	barisan(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	barisan(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%205/OUTPUT/SOAL4.png)
Program Go pada soal nomor 4 bertujuan untuk menampilkan barisan angka dengan pola menurun dari nilai N hingga 1, kemudian dilanjutkan kembali naik dari 1 hingga N menggunakan konsep rekursif. Program menerima input berupa bilangan N dari pengguna, lalu memanggil fungsi rekursif yang pertama-tama mencetak nilai N, kemudian memanggil dirinya sendiri dengan nilai yang dikurangi satu hingga mencapai kondisi berhenti (saat N = 0). Setelah proses rekursi mencapai titik dasar, fungsi akan melanjutkan eksekusi dengan mencetak kembali nilai saat proses “naik” (backtracking), sehingga menghasilkan pola simetris seperti N, N-1, ..., 1, ..., N-1, N. Program ini menunjukkan cara kerja rekursi yang memiliki dua tahap, yaitu saat pemanggilan fungsi (descending) dan saat kembali dari pemanggilan (ascending), sesuai dengan konsep dasar rekursi yang membangun solusi dari pemecahan masalah yang lebih kecil.



## Unguided 

### 5. [Soal]
#### soal5.go

```go
package main

import "fmt"

func ganjil(n int, i int) {
	if i > n {
		return
	}

	if i%2 != 0 {
		fmt.Print(i, " ")
	}

	ganjil(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	ganjil(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%205/OUTPUT/SOAL5.png)
Program Go pada soal nomor 5 bertujuan untuk menampilkan deret bilangan ganjil hingga batas tertentu menggunakan konsep rekursif. Program menerima input berupa sebuah bilangan N dari pengguna, kemudian menggunakan fungsi rekursif yang berjalan dari nilai awal (biasanya 1) hingga N. Pada setiap langkah, fungsi akan mengecek apakah bilangan tersebut merupakan bilangan ganjil dengan kondisi i % 2 != 0, dan jika memenuhi, maka bilangan tersebut akan dicetak. Selanjutnya fungsi memanggil dirinya sendiri dengan nilai yang bertambah hingga mencapai kondisi berhenti saat nilai melebihi N. Program ini menunjukkan bagaimana rekursi dapat digunakan sebagai pengganti perulangan untuk melakukan pengecekan dan penampilan bilangan secara berulang hingga batas tertentu.


## Unguided 

### 6. [Soal]
#### soal6.go

```go
package main

import "fmt"

func pangkat(x int, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan x dan y: ")
	fmt.Scan(&x, &y)

	fmt.Println("Hasil:", pangkat(x, y))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%205/OUTPUT/SOAL6.png)
Program Go pada soal nomor 6 bertujuan untuk menghitung hasil perpangkatan suatu bilangan menggunakan konsep rekursif. Program menerima dua input dari pengguna, yaitu nilai bilangan dasar (x) dan pangkatnya (y), kemudian menggunakan fungsi rekursif untuk menghitung nilai (x^y). Di dalam fungsi tersebut terdapat kondisi dasar (base case) yaitu ketika nilai y = 0, maka hasilnya adalah 1, sedangkan untuk kondisi lainnya (recursive case), fungsi akan mengembalikan hasil perkalian antara x dengan hasil pemanggilan fungsi itu sendiri dengan parameter yang dikurangi satu, yaitu (x * pangkat(x, y-1)). Proses ini akan terus berulang hingga mencapai kondisi dasar, sehingga menghasilkan nilai akhir dari perpangkatan. Program ini menunjukkan bagaimana rekursi bekerja dengan cara memecah perhitungan menjadi bentuk yang lebih sederhana hingga mencapai kondisi berhenti.







