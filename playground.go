package main

import "fmt"

// Struct dengan method
type kaliAja struct {
	a int
	b int
}

// Method untuk menghitung luas
func (j kaliAja) Area() int {
	return j.a * j.b
}
func main() {
	// Struct dasar
	type customers struct {
		name        string
		email       string
		phoneNumber int
		address     string
	}

	// Inisialisasi customer
	x := customers{
		name:        "Jhon",
		email:       "jhony@tph.com",
		phoneNumber: 123,
		address:     "Canada",
	}

	// Inisialisasi struct
	hasil := kaliAja{
		a: 10,
		b: 21,
	}

	// Panggil method Area()
	fmt.Println("Hasil perkalian:", hasil.Area())

	// Print struct customers
	fmt.Println(x)
}
