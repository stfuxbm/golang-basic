package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Mendefinisikan nama penyanyi dan mengubahnya menjadi huruf kecil untuk perbandingan yang lebih mudah
	singer := "Blake shelton"
	singer = strings.ToLower(singer) // Mengonversi nama penyanyi ke huruf kecil agar perbandingan tidak terpengaruh huruf besar

	// Menggunakan switch untuk mengecek tipe penyanyi berdasarkan nama yang telah diubah menjadi huruf kecil
	switch singer {
	case "blake shelton":
		// Jika penyanyi adalah Blake Shelton, output yang diberikan adalah "Old Singer"
		fmt.Println("Old Singer")
	case "morgan wallen":
		// Jika penyanyi adalah Morgan Wallen, output yang diberikan adalah "Modern Country"
		fmt.Println("Modern Country")
	default:
		// Jika penyanyi tidak sesuai dengan yang ada di case, output yang diberikan adalah "not country boy"
		fmt.Println("Not a country boy")
	}

	// Mendapatkan hari ini untuk menentukan waktu dan aktivitas berdasarkan hari
	today := time.Now()
	day := today.Weekday().String() // Mendapatkan nama hari dalam format string

	// Menggunakan switch untuk memberikan pesan yang sesuai berdasarkan hari
	switch day {
	case "Tuesday", "Monday":
		// Jika hari adalah Senin atau Selasa, menunjukkan waktu kerja yang sibuk
		fmt.Println("Waktu kerja sibuk")
	case "Friday":
		// Jika hari adalah Jumat, menunjukkan libur panjang
		fmt.Println("Libur panjang")
	default:
		// Untuk hari selain yang disebutkan, output adalah "kiamat"
		fmt.Println("Kiamat")
	}

}
