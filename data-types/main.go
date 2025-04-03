/*
Tipe Data Dasar: Jenis nilai yang disimpan variabel.
Digunakan untuk klasifikasi data (angka, teks, benar/salah).
Go punya berbagai tipe: numerik (bulat, desimal), boolean, string, dll.
*/
package main

import "fmt"

func main() {
	// Bilangan bulat positif.
	var postiveNumber uint8 = 12 // uint8: 0 s/d 255.

	// Bilangan bulat (bisa negatif).
	var negatifNumber int32 = -21 // int32: contoh.

	// Bilangan desimal (pecahan).
	var decimalNumber = 2.4 // float64: default desimal.

	// Nilai benar atau salah.
	var booleanData bool = false // bool: true/false.

	// Teks.
	var stringData string = "Halo" // string: kumpulan karakter.

	fmt.Printf("Positif: %d, Negatif: %d\n", postiveNumber, negatifNumber)
	fmt.Printf("Desimal: %.2f\n", decimalNumber)
	fmt.Printf("Boolean: %t\n", booleanData)
	fmt.Printf("String: %s\n", stringData)

	// Nilai default.
	var x string // "": string kosong.
	var y int    // 0: integer.
	var z bool   // false: boolean.
	fmt.Printf("Default - String: '%s', Int: %d, Bool: %t\n", x, y, z)
}
