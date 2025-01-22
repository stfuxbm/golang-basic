package main

import "fmt"

/*
Go secara umum memiliki 3 macam operator:
1. Operator Aritmatika
2. Operator Perbandingan
3. Operator Logika

Catatan:
- Operator Aritmatika digunakan untuk operasi matematis.
- Operator Perbandingan membandingkan dua nilai dan menghasilkan true/false.
- Operator Logika digunakan untuk mengombinasikan atau membalik logika boolean.
*/

func main() {
	// Operator Aritmatika
	a := 12
	b := 3

	// Penjelasan:
	// +: Menambahkan dua angka
	// -: Mengurangkan angka kedua dari angka pertama
	// *: Mengalikan dua angka
	// /: Membagi angka pertama dengan angka kedua (integer division jika keduanya bilangan bulat)
	// %: Menghasilkan sisa pembagian angka pertama dengan angka kedua
	fmt.Println("Operator Aritmatika:")
	fmt.Println("Penjumlahan (a + b):", a+b) // Hasilnya: 15
	fmt.Println("Pengurangan (a - b):", a-b) // Hasilnya: 9
	fmt.Println("Perkalian (a * b):", a*b)   // Hasilnya: 36
	fmt.Println("Pembagian (a / b):", a/b)   // Hasilnya: 4
	fmt.Println("Modulus (a % b):", a%b)     // Hasilnya: 0 (karena 12 % 3 = 0)

	// Operator Perbandingan
	// Penjelasan:
	// !=: Mengecek apakah dua nilai tidak sama
	// ==: Mengecek apakah dua nilai sama
	// <: Mengecek apakah nilai pertama lebih kecil dari nilai kedua
	// >: Mengecek apakah nilai pertama lebih besar dari nilai kedua
	// <=: Mengecek apakah nilai pertama lebih kecil atau sama dengan nilai kedua
	// >=: Mengecek apakah nilai pertama lebih besar atau sama dengan nilai kedua
	fmt.Println("\nOperator Perbandingan:")
	fmt.Println("Apakah a tidak sama dengan b (a != b):", a != b)            // True
	fmt.Println("Apakah a sama dengan b (a == b):", a == b)                  // False
	fmt.Println("Apakah a lebih kecil dari b (a < b):", a < b)               // False
	fmt.Println("Apakah a lebih besar dari b (a > b):", a > b)               // True
	fmt.Println("Apakah a lebih kecil atau sama dengan b (a <= b):", a <= b) // False
	fmt.Println("Apakah a lebih besar atau sama dengan b (a >= b):", a >= b) // True

	// Operator Logika
	x := true
	y := false

	// Penjelasan:
	// && (AND): Menghasilkan true jika kedua nilai bernilai true
	// || (OR): Menghasilkan true jika salah satu nilai bernilai true
	// ! (NOT): Membalikkan nilai boolean (true menjadi false dan sebaliknya)
	fmt.Println("\nOperator Logika:")
	fmt.Println("x AND y (x && y):", x && y) // False (karena y adalah false)
	fmt.Println("x OR y (x || y):", x || y)  // True (karena x adalah true)
	fmt.Println("NOT x (!x):", !x)           // False (karena x adalah true, lalu dibalik)

	/*
		Catatan tambahan:
		- Perhatikan prioritas operator. Misalnya, AND (`&&`) memiliki prioritas lebih tinggi daripada OR (`||`).
		- Dalam operasi logika kompleks, gunakan tanda kurung `()` untuk mengatur urutan evaluasi.
	*/
}
