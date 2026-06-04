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
