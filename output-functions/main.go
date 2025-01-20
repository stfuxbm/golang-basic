/*
Fungsi fmt.Print() - digunakan untuk mencetak argumen dengan format default.
    - Tidak ada spasi yang ditambahkan antara argumen.
    - Tidak ada baris baru yang ditambahkan setelah mencetak.

Fungsi fmt.Println() - digunakan untuk mencetak argumen, dengan tambahan:
    - Spasi secara otomatis ditambahkan di antara argumen.
    - Baris baru ditambahkan di akhir output.

Fungsi fmt.Printf() - digunakan untuk mencetak argumen dengan format tertentu:
    - Memungkinkan kita untuk menentukan format output sesuai dengan kebutuhan, misalnya untuk mencetak tipe data atau nilai dari variabel.
    - Menggunakan format specifier (seperti `%v`, `%T`) untuk menyesuaikan tampilan output.
*/

package main

import "fmt"

func main() {

	// Mendeklarasikan beberapa variabel
	var (
		firstName string = "Tom"
		lastName  string = "Hardy"
		isActor   bool   = true
		age       int    = 47
	)

	// Contoh penggunaan fmt.Print()
	fmt.Println("---FUNGSI MENGGUNAKAN Print---")
	fmt.Print(firstName, lastName, isActor, age) // Output: TomHardytrue47 (tanpa spasi antar variabel)

	// Menambahkan spasi dengan print biasa
	print("\n", "\n") // Menambahkan dua baris kosong (spasi 2 kali)

	// Contoh penggunaan fmt.Println()
	fmt.Println("---FUNGSI MENGGUNAKAN PrintLn---")
	const (
		greatActor bool   = true
		address    string = "London"
	)
	fmt.Println(true, "that means is a great actor, and living in", address)
	// Output: true that means is a great actor, and living in London
	// fmt.Println menambahkan spasi secara otomatis antar kata dan baris baru di akhir

	print("\n") // Menambahkan baris kosong lagi

	// Contoh penggunaan fmt.Printf()
	fmt.Println("---FUNGSI MENGGUNAKAN Printf---")
	x := 12
	y := "good"

	// %v untuk mencetak nilai dari variabel
	// %T untuk mencetak tipe data dari variabel
	fmt.Printf("value x adalah: %v dan tipe datanya %T\n", x, x)
	// Output: value x adalah: 12 dan tipe datanya int

	fmt.Printf("value y adalah: %v dan tipe datanya dari %T\n", y, y)
	// Output: value y adalah: good dan tipe datanya dari string
}
