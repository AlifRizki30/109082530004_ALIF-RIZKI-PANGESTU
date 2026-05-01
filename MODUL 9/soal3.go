package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	pemenang := []string{}
	i := 1

	for {
		var a, b int
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&a, &b)

		if a < 0 || b < 0 {
			break
		}

		if a > b {
			pemenang = append(pemenang, klubA)
		} else if b > a {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
		i++
	}

	for i, v := range pemenang {
		fmt.Printf("Hasil %d : %s\n", i+1, v)
	}

	fmt.Println("Pertandingan selesai")
}
