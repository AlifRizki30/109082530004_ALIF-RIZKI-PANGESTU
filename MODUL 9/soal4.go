package main

import "fmt"

const NMAX = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	*n = 0
	for {
		var ch rune
		fmt.Scanf("%c\n", &ch)

		if ch == '.' || *n >= NMAX {
			break
		}
		t[*n] = ch
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var t tabel
	var n int

	fmt.Println("Masukkan teks (akhiri dengan .):")
	isiArray(&t, &n)

	fmt.Print("Teks: ")
	cetakArray(t, n)

	if palindrom(t, n) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}

	balikanArray(&t, n)

	fmt.Print("Reverse teks: ")
	cetakArray(t, n)
}
