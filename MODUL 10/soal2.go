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
