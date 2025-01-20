/*
Konstanta adalah variabel yang memiliki nilai tetap yang tidak dapat diubah setelah dideklarasikan.
Konstanta biasanya dideklarasikan dengan kata kunci `const`.
Syntax untuk mendeklarasikan konstanta adalah:
  const NamaKonstanta tipeData = nilai

Konstanta dapat dideklarasikan baik secara global maupun di dalam fungsi.
Konstanta biasanya ditulis dengan huruf besar semua (untuk membedakan dengan variabel biasa).
Karena sifatnya yang tidak dapat diubah, konstanta hanya bersifat **read-only**.
*/

package main

import "fmt"

// Deklarasi konstanta global
const isHuman bool = true
const location string = "Jakarta"

func main() {

	// Deklarasi konstanta lokal di dalam fungsi
	const isMan bool = true
	const length int = 172

	// Konstanta dengan huruf besar, biasanya untuk konstanta yang penting atau global
	const COUNTRY string = "Indonesia"

	// Konstanta tidak bisa diubah setelah dideklarasikan
	// Kode berikut akan menyebabkan error jika dibuka, karena konstanta bersifat read-only:
	// const piNumber = 3.14
	// piNumber = 10

	fmt.Println(isHuman, isMan, length, location, COUNTRY)

	// Menampilkan pesan bahwa kita akan mendeklarasikan banyak konstanta sekaligus (multi-constant declaration)
	fmt.Println("--- DEKLARASI MULTI VARIABEL KONSTANTA DENGAN BLOK ---")

	// Deklarasi konstanta dalam blok (biasa digunakan untuk banyak konstanta sekaligus)
	// Blok deklarasi lebih mudah dibaca dan sering digunakan untuk banyak variabel konstanta
	const (
		yourName    string  = "Jhon"
		yourAge     int     = 29
		yourNumber  float64 = 3.10
		yourAddress         = "Jakarta" // Tipe data otomatis (string) berdasarkan nilai yang diberikan

	)

	// Menampilkan hasil dari beberapa konstanta sekaligus
	fmt.Println(
		"Ini adalah hasil dari multiple constant variables: ",
		yourName,
		yourAge,
		yourNumber,
		yourAddress,
	)

}
