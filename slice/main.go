/*
Slice: abstraksi di atas array yang ukurannya bisa berubah.
Digunakan untuk koleksi data sejenis yang ukurannya tidak pasti atau perlu diubah.
Slice mereferensikan sebagian atau seluruh array yang mendasarinya.
*/
package main

import "fmt"

func main() {
	// Membuat slice integer.
	x := []int{1, 2, 3, 4, 5} // Tidak perlu ukuran awal.

	// Membuat slice string.
	a := []string{"Tom", "Hardy"}

	fmt.Println("Slice x:", x)
	fmt.Println("Slice a:", a)

	// Mengubah elemen slice.
	x[2] = 10
	a[1] = "Holland"

	fmt.Println("x[2] diubah:", x)
	fmt.Println("a diubah:", a)

	// Menambah elemen ke slice (membuat slice baru).
	x = append(x, 6, 7)
	a = append(a, "Hardy")

	fmt.Println("x setelah append:", x)
	fmt.Println("a setelah append:", a)

	// Memotong slice (mengubah panjang).
	x = x[:4] // Ambil 4 elemen pertama.
	fmt.Println("x dipotong:", x)

	// Menyalin slice.
	originalSlice := []int{100, 200, 300}
	copiedSlice := make([]int, len(originalSlice)) // Buat slice tujuan.
	copy(copiedSlice, originalSlice)               // Salin data.

	fmt.Println("Original:", originalSlice)
	fmt.Println("Copied:", copiedSlice)
}
