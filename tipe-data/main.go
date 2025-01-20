package main

import "fmt"

func main() {

	// Menampilkan informasi tentang tipe data string
	fmt.Println("---TIPE DATA STRING---")
	print("\n")

	// Deklarasi variabel dengan tipe data string
	var firstName string = "Tom Hardy"
	lastName := "Hardy" // Tipe data string dengan pendeklarasian singkat

	// Menampilkan nilai variabel bertipe string
	fmt.Println("string", firstName, lastName)

	print("\n")

	// Menampilkan informasi tentang tipe data boolean
	fmt.Println("--- TIPE DATA BOOLEAN---")
	print("\n")

	/*
		Tipe data boolean hanya memiliki dua nilai: true dan false.
		Jika variabel bertipe boolean tidak didefinisikan nilainya,
		secara default nilainya adalah false.
	*/
	var x bool        // Nilai default: false
	var z bool = true // Menetapkan nilai true pada variabel z
	a := true         // Pendeklarasian singkat, nilai default true

	// Menampilkan nilai variabel bertipe boolean
	fmt.Println(x, z, a)

	print("\n")

	// Menampilkan informasi tentang tipe data integer
	fmt.Println("---TIPE DATA INTEGER---")
	print("\n")

	/*
		Tipe data integer digunakan untuk menyimpan bilangan bulat.
		Integers dibagi menjadi dua kategori:
		- Signed Integers: Dapat menyimpan nilai positif dan negatif.
		- Unsigned Integers: Hanya dapat menyimpan nilai positif.

		Tipe data default untuk integer adalah 0 jika nilai tidak didefinisikan.
		- Jika tipe data integer menggunakan `int`, maka tipe ini disesuaikan dengan sistem operasi (32-bit atau 64-bit).
		Pemilihan tipe data integer perlu diperhatikan, terutama pada aplikasi yang membutuhkan efisiensi memori.
	*/

	// Deklarasi integer dengan nilai positif
	var q int = 45   // Bisa menyimpan nilai positif atau negatif
	var w int32 = 12 // Bisa menyimpan nilai positif atau negatif dengan tipe 32-bit

	// Menampilkan tipe data dari variabel q dan w
	fmt.Printf("tipe data dari %T", q)
	print("\n")
	fmt.Printf("tipe data dari %T", w)

	// Deklarasi unsigned integer (hanya nilai positif)
	var r uint = 600 // Tidak bisa menyimpan nilai negatif
	var t uint = 200 // Tidak bisa menyimpan nilai negatif

	// Menampilkan tipe data dari variabel r dan t
	fmt.Printf("tipe data dari %T", r)
	fmt.Printf("tipe data dari %T", t)

	print("\n")

	// Menampilkan informasi tentang tipe data float
	fmt.Println("---TIPE DATA FLOAT---")
	print("\n")

	/*
		Tipe data float digunakan untuk menyimpan angka desimal.
		Tipe ini bisa menyimpan nilai negatif atau positif dengan desimal.
		Tersedia tipe float32 dan float64. Jika tipe data tidak didefinisikan,
		secara default akan menggunakan float64.
	*/

	// Deklarasi variabel bertipe float32 dan float64
	var y float32 = 3.12 // float32 untuk angka desimal
	var u float64 = -4.5 // float64 untuk angka desimal dengan presisi lebih tinggi

	// Menampilkan nilai dari variabel bertipe float
	fmt.Println(y, u)
}
