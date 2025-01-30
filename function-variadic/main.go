package main

import "fmt"

/*
 Variadic Function dalam Go:
- Bisa menerima banyak nilai dengan tipe yang sama tanpa harus didefinisikan jumlahnya.
- Parameter variadic memiliki sifat yang mirip dengan slice.
- Menggunakan `...` untuk menerima banyak input.

📌 Contoh:
  func namaFungsi(parameter ...tipe) { ... }
  namaFungsi(1, 2, 3, 4, 5) // Bisa menerima banyak angka

Di bawah ini ada beberapa contoh fungsi variadic:
1. `calculateAverage()` -> Menghitung rata-rata dari angka yang diberikan.
2. `showFriendName()` -> Menampilkan daftar nama teman.
3. `findMaxNumber()` -> Mencari angka terbesar dari sekumpulan angka.
*/

// Fungsi untuk menghitung rata-rata angka yang diberikan
func calculateAverage(numbers ...int) float64 {
	// Menjumlahkan semua angka yang diberikan
	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println("Total:", total) // Menampilkan total angka

	// Menghitung rata-rata
	var avg = float64(total) / float64(len(numbers))
	return avg
}

// Fungsi untuk menampilkan daftar teman berdasarkan nama yang diberikan
func showFriendName(names ...string) {
	fmt.Println("List of Friends:")
	for _, name := range names {
		fmt.Println("-", name)
	}
}

// Fungsi untuk mencari angka terbesar dari sekumpulan angka
func findMaxNumber(numbers ...int) int {
	if len(numbers) == 0 {
		fmt.Println("Tidak ada angka yang diberikan")
		return 0
	}

	maxNumber := numbers[0] // Anggap angka pertama sebagai yang terbesar sementara
	for _, number := range numbers {
		if number > maxNumber {
			maxNumber = number
		}
	}
	return maxNumber
}

func main() {
	// Contoh pemanggilan fungsmi calculateAverage dengan angka langsung
	var avg1 = calculateAverage(1, 2, 3)
	fmt.Println("Rata-rata:", avg1) // Output: 2

	//  Contoh pemanggilan calculateAverage dengan slice
	var z = []int{1, 4, 5, 6, 7, 4}
	fmt.Println("Rata-rata dengan slice:", calculateAverage(z...)) // Menggunakan spread operator `...`

	//  Contoh pemanggilan showFriendName dengan daftar teman langsung
	showFriendName("Tom", "Hardy")

	//  Contoh pemanggilan showFriendName dengan slice
	friends := []string{"Jhon", "Wick"}
	showFriendName(friends...) // Menggunakan spread operator `...`

	//  Contoh pemanggilan findMaxNumber untuk mencari angka terbesar
	fmt.Println("Angka terbesar:", findMaxNumber(10, 20, 50, 30, 40))

	//  Contoh pemanggilan findMaxNumber dengan slice angka
	numbers := []int{5, 15, 25, 35, 45}
	fmt.Println("Angka terbesar dari slice:", findMaxNumber(numbers...))
}
