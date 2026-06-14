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
