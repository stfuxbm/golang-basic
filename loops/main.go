package main

import "fmt"

/*
Perulangan (Loop) adalah proses mengulang eksekusi blok kode selama kondisi yang ditentukan masih terpenuhi.
Golang mendukung berbagai jenis perulangan seperti `for`, `nested loop`, dan `for range`.
*/

func main() {
	// Contoh perulangan `for` biasa
	fmt.Println("Perulangan for: mencetak angka 0 hingga 10")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// Nested loop: Mengombinasikan merek dengan jenis perangkat
	devices := [1]string{"Laptop"}
	brands := [2]string{"Dell", "Lenovo"}

	fmt.Println("Nested loop: Mengombinasikan perangkat dan merek")
	for i := 0; i < len(devices); i++ {
		for j := 0; j < len(brands); j++ {
			fmt.Println(devices[i], brands[j])
		}
	}

	// Perulangan menggunakan `for range` dengan slice
	actors := []string{"Tom Hardy", "Heath Ledger"}
	fmt.Println("For range dengan slice: Menampilkan aktor")
	for index, actor := range actors {
		fmt.Printf("Index %d: %s\n", index, actor)
	}

	// Perulangan menggunakan `for range` dengan map
	bestMovies := map[string]string{
		"Aktor": "Tom Holland",
		"Film":  "Spider-Man",
	}
	fmt.Println("For range dengan map: Menampilkan pasangan key-value")
	for key, value := range bestMovies {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Perulangan dengan `break` dan `continue`
	fmt.Println("Perulangan dengan break: Berhenti saat angka mencapai 5")
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break // Berhenti jika i sama dengan 5
		}
		fmt.Println(i)
	}

	fmt.Println("Perulangan dengan continue: Melewati angka genap")
	for j := 0; j <= 10; j++ {
		if j%2 == 0 {
			continue // Melanjutkan ke iterasi berikutnya jika j adalah angka genap sampai batas angka yang di tentukan
		}
		fmt.Println(j)
	}
}
