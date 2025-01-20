package main

import "fmt"

func main() {
	// Tipe Data Numerik Non-Desimal
	var postiveNumber uint8 = 12  // Tipe data uint8 untuk angka positif (range: 0-255)
	var negatifNumber int32 = -21 // Tipe data int32 untuk angka dengan nilai negatif
	fmt.Printf("Angka Positif (uint8): %d, Angka Negatif (int32): %d\n", postiveNumber, negatifNumber)

	// Tipe Data Numerik Desimal
	var decimalNumber = 2.4 // Tipe data float64 secara default untuk nilai desimal untuk spesifikasi tinggi
	fmt.Printf("Angka Desimal (float64): %.2f\n", decimalNumber)

	// Tipe Data bool (Boolean)
	var booleanData bool = false // Tipe data bool hanya bisa true atau false
	fmt.Printf("Data Boolean: %t\n", booleanData)

	// Tipe Data string
	var stringData string = "Halo" // Tipe data string untuk menyimpan teks
	fmt.Printf("Data String: %s\n", stringData)

	// Nilai nil & Zero Value
	var x string // string kosong "" adalah nilai default untuk tipe string
	var y int    // 0 adalah nilai default untuk tipe int
	var z bool   // false adalah nilai default untuk tipe bool
	fmt.Printf("Nilai Default - String: '%s', Int: %d, Bool: %t\n", x, y, z)

	// Tipe Data Numerik Lebih Besar
	var bigInt int64 = 9223372036854775807    // Tipe data int64 untuk nilai lebih besar
	var bigUint uint64 = 18446744073709551615 // Tipe data uint64 untuk nilai lebih besar
	fmt.Printf("Angka Besar (int64): %d, Angka Besar Tanpa Tanda (uint64): %d\n", bigInt, bigUint)

	// Tipe Data Desimal Lainnya (float32)
	var decimal32 float32 = 3.14 // Tipe data float32 untuk angka desimal dengan presisi lebih rendah untuk memori kecil
	fmt.Printf("Angka Desimal (float32): %.2f\n", decimal32)

	// Tipe Data Complex
	var complexNum complex64 = 1 + 2i // Tipe data complex dengan nilai real dan imajiner
	fmt.Printf("Angka Kompleks (complex64): %v\n", complexNum)

	// Tipe Data Pointer
	var ptr *int = &y // Pointer ke variabel y
	fmt.Printf("Pointer ke y: %p\n", ptr)
}
