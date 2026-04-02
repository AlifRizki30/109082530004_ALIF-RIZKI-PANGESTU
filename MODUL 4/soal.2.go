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
