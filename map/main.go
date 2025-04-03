/*
Variabel: tempat menyimpan data dengan nama.
Digunakan untuk menyimpan nilai yang bisa berubah selama program berjalan.
Deklarasi menggunakan 'var' atau ':=', tipe data bisa eksplisit atau otomatis.
*/
package main

import "fmt"

func main() {
	// Deklarasi tanpa tipe eksplisit (inferensi).
	var a = "hello" // 'a' otomatis jadi string.

	// Deklarasi dengan tipe eksplisit.
	var b string = "world"

	// Deklarasi singkat (hanya di dalam fungsi).
	c := 10 // 'c' otomatis jadi integer.

	// Deklarasi banyak variabel dalam satu blok.
	var (
		d = true
		e int
	)

	fmt.Println(a, b, c, d, e)
}
