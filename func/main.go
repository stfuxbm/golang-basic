/*
Function: blok kode yang dapat dipanggil berkali-kali.
Digunakan untuk mengorganisir kode, menghindari pengulangan, dan membuat program modular.
Nama fungsi umumnya menggunakan camelCase.
Parameter adalah variabel input fungsi.
*/
package main

import "fmt"

func main() {
	basicFunc()                                    // Panggil fungsi tanpa parameter/return.
	getUserName("Halo", "Tom Hardy")               // Panggil fungsi dengan parameter.
	fmt.Println(getResult(1, 4, 6))                // Panggil fungsi dengan return value.
	getResultWithVariable(1, 5, 6)                 // Panggil fungsi tanpa return, cetak di dalam.
	fmt.Println(calculateDivisionAndRemainder(10)) // Panggil fungsi dengan multiple return values.
}

// Fungsi dasar tanpa parameter/return.
func basicFunc() {
	fmt.Println("Hallo")
}

// Fungsi dengan parameter (input).
func getUserName(message string, userName string) {
	fmt.Println(message, userName)
}

// Fungsi dengan return value (output).
func getResult(a int, b int, c int) int {
	return a + b/c
}

// Fungsi dengan named return value (output bernama).
func getResultWithVariable(a int, b int, c int) {
	res := a*b + c
	fmt.Println(res)
}

// Fungsi dengan multiple return values (banyak output).
func calculateDivisionAndRemainder(x int) (quotient, remainder int) {
	quotient = x / 3
	remainder = x % 3
	return // Return nilai quotient dan remainder.
}
