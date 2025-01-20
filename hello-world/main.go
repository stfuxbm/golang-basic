package main

import "fmt" // Paket untuk operasi input/output seperti Println

func main() {
	// Deklarasi variabel tanpa tipe data eksplisit
	var a = "hello world"
	// Penjelasan: Variabel `a` dideklarasikan menggunakan `var` tanpa tipe eksplisit  type inference (inferensi tipe).
	// Golang otomatis mengenali tipe data sebagai `string` karena nilai awalnya berupa teks.

	// Deklarasi variabel dengan tipe data string secara eksplisit
	var x string = "hello world"
	// Penjelasan: Variabel `x` dideklarasikan dengan tipe `string` secara eksplisit
	// dan langsung diberi nilai "hello world".

	// Deklarasi variabel dengan shorthand assignment
	y := "hello World"
	// Penjelasan: Variabel `y` dideklarasikan menggunakan shorthand `:=`.
	// Tipe data ditentukan secara otomatis berdasarkan nilai yang diberikan.

	// Deklarasi variabel menggunakan blok `var`
	var (
		z = "hello world"
	)
	// Penjelasan: Variabel `z` dideklarasikan dalam blok `var`.
	// Tidak disebutkan tipe datanya secara eksplisit, tetapi Golang secara otomatis
	// mengenali tipe data sebagai `string` karena nilai yang diberikan.

	// Output: Mencetak semua variabel ke konsol
	fmt.Println(a, x, y, z)
	// Penjelasan: `fmt.Println` mencetak nilai `a`, `x`, `y`, dan `z` ke konsol.
	// Output akan dipisahkan secara otomatis oleh spasi.
}
