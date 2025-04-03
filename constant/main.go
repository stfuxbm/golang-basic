/*
Konstanta: variabel dengan nilai tetap setelah deklarasi.
Digunakan untuk nilai yang tidak boleh berubah sepanjang program,
seperti nilai matematika (pi), konfigurasi, atau string literal yang sering dipakai.
Deklarasikan dengan 'const'.
*/
package main

import "fmt"

func main() {
	// Konstanta integer.
	const x int = 3

	// Konstanta string.
	const y string = "halo"

	// Konstanta boolean.
	const z bool = true

	fmt.Println(x, y, z)

	// Beberapa konstanta dalam satu blok.
	const (
		a string = "hai"
		b int    = 2
		c bool   = false
		d        = 3.14 // float otomatis
	)

	fmt.Println(a, b, c)

	// Beberapa konstanta dalam satu baris (tipe otomatis).
	const q, w, r = 1, 2, 3 // Semua int

	fmt.Println(q, w, r)
}
