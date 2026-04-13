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
