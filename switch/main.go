/*
Switch Statement: alternatif untuk if-else if berantai.
Digunakan untuk memilih satu blok kode untuk dieksekusi berdasarkan nilai variabel.
Berguna saat membandingkan satu variabel dengan banyak kemungkinan nilai.
*/
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Contoh switch pada string.
	singer := "Blake shelton"
	singerLower := strings.ToLower(singer) // Ubah ke huruf kecil.

	switch singerLower {
	case "blake shelton": // Jika singerLower adalah "blake shelton".
		fmt.Println("Old Singer")
	case "morgan wallen": // Jika singerLower adalah "morgan wallen".
		fmt.Println("Modern Country")
	default: // Jika tidak ada case yang cocok.
		fmt.Println("Not a country boy")
	}

	// Contoh switch pada tipe lain (hari).
	today := time.Now()
	day := today.Weekday().String() // Dapatkan nama hari.

	switch day {
	case "Tuesday", "Monday": // Jika hari Senin atau Selasa.
		fmt.Println("Waktu kerja sibuk")
	case "Friday": // Jika hari Jumat.
		fmt.Println("Libur panjang")
	default: // Untuk hari lainnya.
		fmt.Println("Kiamat")
	}
}
