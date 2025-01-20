package main

import "fmt"

func main() {
	/*
		Konstanta adalah jenis variabel yang nilainya tidak bisa diubah setelah dideklarasikan.
		Inisialisasi nilai konstanta hanya dilakukan sekali saja di awal, setelah itu variabel tidak bisa diubah nilainya.
		Konstanta dalam Go tidak bisa dideklarasikan dengan shorthand assignment `:=`.
		Konstanta hanya bisa dideklarasikan dengan `const`.
	*/

	// Deklarasi konstanta dengan tipe data eksplisit
	const x int = 3         // Konstanta dengan tipe eksplisit int
	const y string = "halo" // Konstanta dengan tipe eksplisit string
	const z bool = true     // Konstanta dengan tipe eksplisit bool

	// Menampilkan nilai-nilai konstanta
	fmt.Println(x, y, z)

	// Deklarasi Multi Konstanta dalam satu blok
	const (
		a string = "hai" // Konstanta string
		b int    = 2     // Konstanta int
		c bool   = false // Konstanta bool
		d        = 3.14  // Konstanta float64 secara otomatis ditentukan oleh Go
	)

	// Menampilkan nilai-nilai konstanta
	fmt.Println(a, b, c)

	// Deklarasi Multiple Konstanta dalam satu baris dengan tipe otomatis
	const q, w, r = 1, 2, 3 // Semua konstanta ini akan bertipe `int` secara otomatis

	// Menampilkan nilai-nilai konstanta yang dideklarasikan dalam satu baris
	fmt.Println(q, w, r)
}
