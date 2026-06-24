# <h1 align="center">sLaporan Praktikum Modul 3 - ... </h1>
<p align="center">[ALIF RIZKI PANGESTU] - [109082530004]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}
func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	p1 := faktorial(a) / faktorial(a-c)
	k1 := faktorial(a) / (faktorial(c) * faktorial(a-c))
	p2 := faktorial(b) / faktorial(b-d)
	k2 := faktorial(b) / (faktorial(d) * faktorial(b-d))
	fmt.Println(p1, k1)
	fmt.Println(p2, k2)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%203/OUTPUT/soal%201%20modul%203.png)
Program Go ini intinya adalah sebuah kalkulator yang bisa menghitung permutasi (P) dan kombinasi (K) dari dua pasang angka sekaligus. Cara kerjanya simpel — program minta kita masukkan 4 angka (a, b, c, d), lalu dia otomatis hitung berapa banyak cara kita bisa menyusun atau memilih sejumlah objek menggunakan rumus faktorial di balik layar. Misalnya waktu dimasukkan angka 5, 10, 3, 10 hasilnya permutasi dan kombinasi dari pasangan pertama adalah 60 dan 10, sedangkan percobaan kedua dengan angka 8, 0, 2, 0 menghasilkan 56 dan 1 — semua dihitung otomatis tanpa kita perlu repot hitung manual satu-satu.
## Unguided 

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func f(x int) int { return x * x }
func g(x int) int { return x - 2 }
func h(x int) int { return x + 1 }

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%203/OUTPUT/soal%202%20modul%203.png)
Program Go ini pada dasarnya adalah latihan fungsi bersarang (nested function), di mana ada tiga fungsi matematika sederhana — f(x) mengkuadratkan angka (x × x), g(x) mengurangi angka dengan 2 (x - 2), dan h(x) menambah angka dengan 1 (x + 1) — lalu ketiganya digabung-gabungkan secara berlapis. Jadi ketika kita masukkan 3 angka (a, b, c), program tidak hanya pakai satu fungsi, tapi memasukkan hasil satu fungsi ke fungsi lain secara berantai, misalnya f(g(h(a))) artinya angka a dulu ditambah 1, lalu dikurangi 2, lalu dikuadratkan. Hasilnya beda-beda tergantung urutan fungsinya, dan ini membuktikan bahwa urutan fungsi itu sangat berpengaruh terhadap hasil akhir — sama seperti di matematika, f(g(x)) belum tentu sama hasilnya dengan g(f(x)).


## Unguided 

### 3. [Soal]
#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}
func diDalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}
func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)
	dalam1 := diDalam(cx1, cy1, r1, x, y)
	dalam2 := diDalam(cx2, cy2, r2, x, y)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AlifRizki30/109082530004_ALIF-RIZKI-PANGESTU/blob/main/MODUL%203/OUTPUT/soal%203%20modul%203.png)
Program Go ini ibarat seorang penjaga yang bertugas mengecek apakah sebuah titik berada di dalam satu atau dua lingkaran sekaligus. Caranya, kita cukup masukkan pusat dan jari-jari dari dua lingkaran, lalu masukkan koordinat titik yang ingin dicek — program akan otomatis menghitung jarak titik tersebut ke masing-masing pusat lingkaran menggunakan rumus jarak (Pythagoras), dan kalau jaraknya lebih kecil atau sama dengan jari-jarinya berarti titik itu ada di dalam lingkaran tersebut. Hasilnya pun langsung jelas — program akan bilang apakah titik itu ada "di dalam lingkaran 1 dan 2" (masuk keduanya), "di dalam lingkaran 1" saja, "di dalam lingkaran 2" saja, atau "di luar kedua lingkaran" — jadi sangat praktis untuk menentukan posisi suatu titik terhadap beberapa lingkaran tanpa perlu hitung manual.









