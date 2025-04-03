/*
Variabel: tempat menyimpan data dengan nama tertentu.
Digunakan untuk menyimpan nilai yang dapat diubah selama program berjalan.
Go memiliki beberapa cara deklarasi variabel.
*/
package main

import "fmt" // Untuk fungsi input/output.

func main() {
	// Deklarasi variabel tanpa tipe eksplisit (type inference).
	var a = "hello world"
	// `var a = ...`: Deklarasi variabel 'a'.
	// Go otomatis menentukan tipe data (string) dari nilai awal.

	// Deklarasi variabel dengan tipe string eksplisit.
	var x string = "hello world"
	// `var x string = ...`: Deklarasi variabel 'x' dengan tipe string.

	// Deklarasi variabel dengan shorthand assignment (hanya di dalam fungsi).
	y := "hello World"
	// `y := ...`: Deklarasi dan inisialisasi singkat.
	// Tipe data otomatis ditentukan (string).

	// Deklarasi variabel menggunakan blok `var`.
	var (
		z = "hello world"
	)
	// `var (...)`: Deklarasi beberapa variabel dalam satu blok.
	// Tipe data otomatis ditentukan (string).

	// Cetak nilai semua variabel.
	fmt.Println(a, x, y, z)
	// `fmt.Println(...)`: Mencetak nilai variabel ke konsol, dipisahkan spasi.
}
