package main

import "fmt"

/*
Pointer adalah reference atau alamat memori.
Variabel pointer berarti variabel yang berisi alamat memori suatu nilai.
*/

func main() {
	// Membuat variabel x dengan nilai 4
	x := 4

	// Mendeklarasikan pointer z yang menyimpan alamat memori dari x
	z := &x

	// Menampilkan nilai dari x, alamat pointer z, dan dereference pointer z
	fmt.Println("Value of x:", x)             // Output: 4
	fmt.Println("Pointer (Alamat x):", z)     // Output: alamat memori dari x (contoh: 0xc0000120a0)
	fmt.Println("Dereference Pointer z:", *z) // Output: 4, karena *z mengambil nilai dari alamat memori x

	// Mengubah nilai x
	x = 2
	fmt.Println("\nSetelah mengubah nilai x:")
	fmt.Println("Value of x:", x)             // Output: 2
	fmt.Println("Pointer (Alamat x) z:", z)   // Alamat pointer tetap sama, hanya nilai x yang berubah
	fmt.Println("Dereference Pointer z:", *z) // Output: 2, karena *z mengambil nilai terbaru dari x

	// Menggunakan pointer dalam parameter fungsi
	a := 12
	fmt.Println("\nSebelum perubahan, nilai a:", a) // Output: 12
	// Memanggil fungsi untuk mengubah nilai a melalui pointer
	changesPointer(&a, 1)
	fmt.Println("Setelah perubahan, nilai a:", a) // Output: 1
}

// Fungsi untuk mengubah nilai variabel yang ditunjuk oleh pointer
func changesPointer(original *int, value int) {
	*original = value // Mengubah nilai variabel yang ditunjuk oleh pointer karena original adalah pointer, bukan nilai
}
