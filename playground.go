package main

import (
	"fmt"
	"strings"
)

/*
Fungsi adalah sekumpulan blok kode yang dibungkus dengan nama tertentu.
Penerapan fungsi yang tepat akan menjadikan kode lebih modular dan juga DRY (Don't Repeat Yourself).
*/

func main() {

	var brands = []string{
		"Lenovo",
		"Samsuung",
	}
	// Memanggil fungsi untuk menampilkan data dengan motode void function
	showBrands("hello", brands)

	// memanggil fungsi untuk menampilkan data dengan metode return value

}

// Penerapan Fungsi dengan simple atau bias di sebut void function
func showBrands(message string, listBrands []string) {
	// Menggabungkan elemen-elemen dalam slice dengan spasi
	var brandName = strings.Join(listBrands, " ")
	fmt.Println(message, brandName)
}

// Penerapan Fungsi dengan Return Value
