package main

import "fmt"

func printPattern(n int, current int) {
	if current > n {
		return
	}

	// cetak bintang
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
