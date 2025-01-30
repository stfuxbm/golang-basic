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

	// Contoh void function dengan parameter slice string (nama)
	var name = []string{
		"Alice",
		"Bob",
	}
	callName("Hello", name) // Menampilkan pesan dengan daftar nama

	// Contoh void function dengan parameter angka
	var number = []int{
		1, 2,
	}
	addTwoNumbers(number[0], number[1]) // Menambahkan dua angka
	sumOfNumbers(number)                // Menjumlahkan seluruh angka dalam slice

	// Contoh fungsi dengan return value
	res := returnValues(1, 2) // Menambahkan dua angka dan mengembalikan hasilnya
	fmt.Println(res)          // Menampilkan hasil penjumlahan
}

// Fungsi void yang menampilkan pesan dan menggabungkan nama dalam slice menjadi string
func callName(message string, userName []string) {
	// Menggabungkan elemen-elemen dalam slice dengan koma sebagai pemisah
	joinNames := strings.Join(userName, ",")
	fmt.Println(message, joinNames)
}

// Fungsi void yang menambahkan dua angka
func addTwoNumbers(a int, b int) {
	calculateNumber := a + b
	fmt.Println(calculateNumber) // Menampilkan hasil penjumlahan
}

// Fungsi void yang menjumlahkan seluruh angka dalam slice
func sumOfNumbers(numbers []int) {
	total := 0
	// Iterasi untuk menjumlahkan seluruh angka dalam slice
	for _, number := range numbers { // _, untuk meenghilangkan index dari elemet nya
		total += number
	}
	fmt.Println(total) // Menampilkan total penjumlahan
}

// Fungsi dengan return value yang menjumlahkan dua angka
func returnValues(a, b int) int {
	return a + b // Mengembalikan hasil penjumlahan
}
