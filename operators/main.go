/*
Operator: simbol untuk melakukan operasi pada nilai (operand).
Digunakan untuk manipulasi data (aritmatika), perbandingan, dan logika.
Go punya operator aritmatika, perbandingan, dan logika.
*/
package main

import "fmt"

func main() {
	a := 12
	b := 3
	x := true
	y := false

	// Operator Aritmatika: operasi matematika.
	fmt.Println("Aritmatika:")
	fmt.Println("a + b =", a+b) // Penjumlahan.
	fmt.Println("a - b =", a-b) // Pengurangan.
	fmt.Println("a * b =", a*b) // Perkalian.
	fmt.Println("a / b =", a/b) // Pembagian (integer jika bulat).
	fmt.Println("a % b =", a%b) // Sisa bagi (modulus).

	// Operator Perbandingan: menghasilkan true atau false.
	fmt.Println("\nPerbandingan:")
	fmt.Println("a != b =", a != b) // Tidak sama dengan.
	fmt.Println("a == b =", a == b) // Sama dengan.
	fmt.Println("a < b =", a < b)   // Lebih kecil dari.
	fmt.Println("a > b =", a > b)   // Lebih besar dari.
	fmt.Println("a <= b =", a <= b) // Lebih kecil atau sama dengan.
	fmt.Println("a >= b =", a >= b) // Lebih besar atau sama dengan.

	// Operator Logika: mengkombinasikan/membalik boolean.
	fmt.Println("\nLogika:")
	fmt.Println("x && y =", x && y) // AND (keduanya true).
	fmt.Println("x || y =", x || y) // OR (salah satu true).
	fmt.Println("!x =", !x)         // NOT (kebalikan).
}
