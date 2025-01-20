// Go adalah bahasa pemrograman yang bersifat compiled, seperti halnya file exe di Windows.
// Misalnya, ketika kita menjalankan perintah 'go build' pada Windows, hasilnya adalah file .exe yang sesuai dengan sistem operasi Windows.
// Begitu juga untuk sistem operasi lainnya, seperti macOS, yang menghasilkan file dengan ekstensi .dmg.
// Oleh karena itu, hasil akhir dari pengembangan aplikasi Go di production adalah file yang sudah dikompilasi.

package main // Deklarasi package 'main', yang wajib ada di program Go

import "fmt" // Mengimpor package 'fmt' untuk memudahkan input-output (I/O), seperti mencetak output ke konsol

// Fungsi utama 'main' adalah tempat di mana program dieksekusi
func main() {
	// Pernyataan yang akan dieksekusi untuk mencetak output ke konsol
	fmt.Println("hello world")          // Menampilkan "hello world"
	fmt.Println("my name is Braiyen")   // Menampilkan "my name is Braiyen"
	fmt.Println("i am learning golang") // Menampilkan "i am learning golang"
	fmt.Println("hell ya !")            // Menampilkan "hell ya !"
}
