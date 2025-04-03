/*
Perulangan (Loop): eksekusi blok kode berulang kali.
Digunakan untuk melakukan tindakan yang sama atau serupa beberapa kali.
Go mendukung 'for', 'nested for', dan 'for range'.
*/
package main

import "fmt"

func main() {
	// Perulangan for biasa.
	fmt.Println("for: 0 sampai 10")
	for i := 0; i <= 10; i++ { // Inisialisasi, kondisi, increment.
		fmt.Println(i)
	}

	// Nested loop (loop di dalam loop).
	devices := [1]string{"Laptop"}
	brands := [2]string{"Dell", "Lenovo"}
	fmt.Println("Nested loop: perangkat dan merek")
	for i := 0; i < len(devices); i++ { // Loop untuk perangkat.
		for j := 0; j < len(brands); j++ { // Loop untuk merek.
			fmt.Println(devices[i], brands[j])
		}
	}

	// Perulangan for range untuk slice.
	actors := []string{"Tom Hardy", "Heath Ledger"}
	fmt.Println("for range slice: aktor")
	for index, actor := range actors { // Dapatkan indeks dan nilai.
		fmt.Printf("Index %d: %s\n", index, actor)
	}

	// Perulangan for range untuk map.
	bestMovies := map[string]string{"Aktor": "Tom Holland", "Film": "Spider-Man"}
	fmt.Println("for range map: key-value")
	for key, value := range bestMovies { // Dapatkan key dan value.
		fmt.Printf("%s: %s\n", key, value)
	}

	// Perulangan dengan break (keluar dari loop).
	fmt.Println("for dengan break: berhenti di 5")
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break // Keluar dari loop jika i == 5.
		}
		fmt.Println(i)
	}

	// Perulangan dengan continue (lewati iterasi saat ini).
	fmt.Println("for dengan continue: lewati genap")
	for j := 0; j <= 10; j++ {
		if j%2 == 0 {
			continue // Lanjutkan ke iterasi berikutnya jika j genap.
		}
		fmt.Println(j)
	}
}
