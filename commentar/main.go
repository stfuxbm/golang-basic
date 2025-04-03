/*
Komentar dalam Go digunakan untuk memberikan penjelasan pada kode.
Komentar tidak dieksekusi oleh kompiler.
Ada dua jenis komentar: inline (satu baris) dan multiline (lebih dari satu baris).
Komentar penting untuk membuat kode lebih mudah dipahami.
*/
package main

import "fmt"

func main() {

	// komentar inline: penjelasan singkat di akhir baris kode.
	a := "hello" // variabel 'a' menyimpan string "hello".

	/*
	   ini adalah komentar multiline:
	   digunakan untuk penjelasan yang lebih detail
	   yang membutuhkan beberapa baris.
	*/
	var b int = 7 // variabel 'b' bertipe integer dan bernilai 7.

	fmt.Println(a, b) // mencetak nilai variabel 'a' dan 'b' ke konsol.
}
