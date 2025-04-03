/*
Fungsi anonim: fungsi tanpa nama.
Bisa langsung dieksekusi atau disimpan di variabel.
Berguna untuk fungsi kecil, callback, goroutine.
*/
package main

import "fmt"

func main() {

	// Fungsi anonim sederhana, langsung jalan.
	func() { // Definisikan.
		fmt.Println("Ini contoh dasar") // Lakukan aksi.
	}() // Langsung jalankan.

	fmt.Println() // Spasi output.

	// Fungsi anonim dengan parameter, langsung jalan.
	func(name string) { // Terima nama.
		fmt.Println("Hai", name) // Sapa nama.
	}("Jhon") // Panggil dengan nama.

	fmt.Println() // Spasi output.

	// Fungsi anonim disimpan di variabel.
	great := func(country string) string { // Terima negara, kembalikan sapaan.
		return "Hai, " + country // Buat sapaan.
	}
	fmt.Println(great("Jhon")) // Panggil via variabel.
}
