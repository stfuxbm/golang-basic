package main

import "fmt"

/*
Operator penugasan digunakan untuk menetapkan nilai ke dalam sebuah variabel
atau mempersingkat operasi matematika pada variabel itu sendiri.

Contoh operator penugasan:
x = 10  // Menetapkan nilai 10 ke variabel x
x += 10 // Sama dengan x = x + 10
x -= 10 // Sama dengan x = x - 10
x *= 10 // Sama dengan x = x * 10
x /= 10 // Sama dengan x = x / 10
x %= 10 // Sama dengan x = x % 10
*/

func main() {
	var (
		x int = 2 // Inisialisasi variabel x dengan nilai 2
		y int = 3 // Inisialisasi variabel y dengan nilai 3
		z int = 5 // Inisialisasi variabel z dengan nilai 5
	)

	// Operasi penugasan
	x += 3 // x = x + 3, hasilnya x = 2 + 3 = 5
	y -= 2 // y = y - 2, hasilnya y = 3 - 2 = 1
	z %= 5 // z = z % 5, hasilnya z = 5 % 5 = 0

	// Cetak hasil operasi
	fmt.Println(x) // Output: 5
	fmt.Println(y) // Output: 1
	fmt.Println(z) // Output: 0
}
