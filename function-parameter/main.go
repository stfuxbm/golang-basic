package main

import "fmt"

/*
 Menggunakan Fungsi sebagai Parameter dalam Go:
- Memungkinkan kita meneruskan logika fungsi ke fungsi lain.
- Memisahkan logika dari eksekusi untuk fleksibilitas lebih besar.
- Cocok untuk callback functions atau fungsi reusable.

📌 Contoh:
1. `calculate()` -> Menerima fungsi perhitungan sebagai parameter.
2. `printNames()` -> Menerima fungsi untuk menampilkan daftar nama.
3. `processNumbers()` -> Menerima fungsi yang bekerja pada sekumpulan angka.
*/

// Fungsi Higher-Order yang menerima fungsi sebagai parameter
func calculate(operation func(numbers ...int) float64, numbers ...int) float64 {
	return operation(numbers...)
}

// Fungsi Higher-Order untuk mencetak daftar nama
func printNames(printer func(names ...string), names ...string) {
	printer(names...)
}

// Fungsi Higher-Order untuk memproses angka dengan fungsi tertentu
func processNumbers(operation func(numbers ...int) int, numbers ...int) int {
	return operation(numbers...)
}

func main() {
	//  Closure untuk menghitung rata-rata
	averageFunc := func(numbers ...int) float64 {
		total := 0
		for _, number := range numbers {
			total += number
		}
		fmt.Println("Total:", total) // Menampilkan total angka

		// Menghitung rata-rata
		return float64(total) / float64(len(numbers))
	}

	//  Closure untuk menampilkan daftar nama teman
	printNamesFunc := func(names ...string) {
		fmt.Println("List of Friends:")
		for _, name := range names {
			fmt.Println("-", name)
		}
	}

	//  Closure untuk mencari angka terbesar
	findMaxFunc := func(numbers ...int) int {
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

	//  Contoh pemanggilan calculate() dengan fungsi rata-rata
	var avg1 = calculate(averageFunc, 1, 2, 3)
	fmt.Println("Rata-rata:", avg1) // Output: 2

	//  Contoh pemanggilan calculate() dengan slice angka
	var z = []int{1, 4, 5, 6, 7, 4}
	fmt.Println("Rata-rata dengan slice:", calculate(averageFunc, z...))

	//  Contoh pemanggilan printNames() untuk menampilkan daftar nama
	printNames(printNamesFunc, "Tom", "Hardy")

	//  Contoh pemanggilan printNames() dengan slice
	friends := []string{"Jhon", "Wick"}
	printNames(printNamesFunc, friends...)

	//  Contoh pemanggilan processNumbers() untuk mencari angka terbesar
	fmt.Println("Angka terbesar:", processNumbers(findMaxFunc, 10, 20, 50, 30, 40))

	//  Contoh pemanggilan processNumbers() dengan slice angka
	numbers := []int{5, 15, 25, 35, 45}
	fmt.Println("Angka terbesar dari slice:", processNumbers(findMaxFunc, numbers...))
}
