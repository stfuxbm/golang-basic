package main

import "fmt"

/*
Operator untuk melakukan operasi umum matematika:
+  = Penjumlahan
-  = Pengurangan
*  = Perkalian
/  = Pembagian
%  = Modulus (Sisa hasil pembagian)
*/

func main() {

	fmt.Println("---OPERATOR DASAR---")

	var (
		x = 10
		y = 3
	)

	// Menampilkan hasil operasi dasar matematika
	fmt.Println(x + y) // Penjumlahan
	fmt.Println(x - y) // Pengurangan
	fmt.Println(x * y) // Perkalian
	fmt.Println(x / y) // Pembagian
	fmt.Println(x % y) // Modulus (Sisa dari hasil pembagian)

	print("\n")
	var a, b = 100, 432
	// Menampilkan sisa pembagian antara a dan b
	fmt.Println(a % b)

	// Menyimpan hasil perkalian dalam variabel hasil
	var hasil = a * y
	println("Hasil perkalian a * y adalah: ", hasil)

	print("\n")
	fmt.Println("---OPERATOR INCREMENT DAN DECREMENT---") // Menaikkan nilai variabel sebesar 1 dan menurunkan variabel sebesar 1

	/*
		Increment (penambahan 1) = i++ sama dengan i = i + 1
		Decrement (pengurangan 1) = i-- sama dengan i = i - 1
	*/

	// Increment
	i := 10
	i++ // Sama saja dengan i = i + 1

	// Decrement
	j := 20
	j-- // Sama saja dengan j = j - 1

	// Menampilkan hasil setelah increment dan decrement
	fmt.Println("Hasil dari increment (i = 10 + 1) adalah ", i)
	fmt.Println("Hasil dari decrement (j = 20 - 1) adalah ", j)
}
