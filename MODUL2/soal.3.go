package main

import "fmt"
func main() {
	var gram int
	var Kg, Sisa int
	var ongkosKg, ongkosSisa, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gram)

	Kg = gram / 1000
	Sisa = gram % 1000
	ongkosKg = Kg * 10000

	if Kg > 10 {
		ongkosSisa = 0
	} else {
		if Sisa >= 500 {
			ongkosSisa = Sisa * 5
		} else {
			ongkosSisa = Sisa * 15
		}
	}
	total = ongkosKg + ongkosSisa

	fmt.Println("Detail berat:", Kg, "Kg+", Sisa, "gr")
	fmt.Println("Detail ongkos: Rp. ", ongkosKg, "+ Rp. ", ongkosSisa)
	fmt.Println("Total ongkos Rp.", total)
}