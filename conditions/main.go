package main

import "fmt"

/*
Kondisi dalam Golang:
1. If-Else: Digunakan untuk seleksi kondisi sederhana.
2. Switch: Alternatif dari if-else ketika ada banyak kasus.
3. Go tidak mendukung ternary operator seperti (a > b ? "yes" : "no").
*/

func main() {
	// **If-Else**
	x := 12
	if x > 12 {
		fmt.Println("Good")
	} else if x == 12 { // Jika sama dengan 12
		fmt.Println("Medium")
	} else {
		fmt.Println("Poor")
	}

	// **Variabel Temporary pada If-Else**
	z := 10
	if p := z / 100; p > 0 { // Variabel p hanya berlaku di dalam blok ini
		fmt.Println("Hasilnya adalah", z)
	} else {
		fmt.Println("Nilai z tidak memenuhi syarat")
	}

	// **Switch Case**
	o := 34
	switch o {
	case 1:
		fmt.Println("Nilai adalah 1")
	case 2:
		fmt.Println("Nilai adalah 2")
	case 34:
		fmt.Println("Nilai adalah 34")
	default:
		fmt.Println("Nilai tidak dikenal")
	}

	// **Switch Case Multi-Kondisi**
	e := 5.0
	switch e {
	case 1:
		fmt.Println("Nilai adalah 1")
	case 2, 5, 8: // Salah satu dari nilai ini
		fmt.Println("Nilai cocok dengan multi-kondisi")
	default:
		fmt.Println("Nilai tidak sesuai dengan kasus mana pun")
	}

	// **Nested Conditions**
	point := 10
	if point > 7 {
		switch point {
		case 10:
			fmt.Println("Perfect!")
		default:
			fmt.Println("Nice!")
		}
	} else { // Jika point <= 7
		if point == 5 {
			fmt.Println("Not bad")
		} else if point == 3 {
			fmt.Println("Keep trying")
		} else {
			fmt.Println("You can do it!")
			if point == 0 {
				fmt.Println("Try harder!")
			}
		}
	}
}
