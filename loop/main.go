package main

import "fmt"

/*
Golang memiliki `for` sebagai satu-satunya konstruksi perulangan.
Perulangan dapat dikombinasikan dengan `break` dan `continue`.
*/
func main() {

	// Perulangan sederhana menggunakan for
	for i := 0; i <= 12; i++ {
		fmt.Println(i)
	}

	// Perulangan dengan break & continue
	for x := 0; x <= 12; x++ {
		if x == 5 {
			continue // Lewati iterasi jika x sama dengan 5
		}
		if x == 10 {
			break // Hentikan perulangan jika x sama dengan 10
		}
		fmt.Println("break and continue:", x)
	}

	// Perulangan bersarang: mencetak angka dalam bentuk pola
	for z := 0; z < 5; z++ {
		for q := z; q < 5; q++ {
			fmt.Print(q, " ") // Cetak angka dengan spasi
		}
		fmt.Println() // Pindah ke baris baru setelah setiap iterasi z
	}
}
