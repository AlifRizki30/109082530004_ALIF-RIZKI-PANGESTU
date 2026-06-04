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
