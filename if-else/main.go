/*
If Statement: untuk eksekusi kondisional blok kode.
Digunakan saat ingin menjalankan kode tertentu hanya jika suatu kondisi terpenuhi.
'if' untuk kondisi pertama, 'else if' untuk kondisi lanjutan, 'else' untuk kondisi terakhir.
Operator logika: `==` (sama dengan), `!=` (tidak sama dengan), `>` (lebih besar), `<` (lebih kecil),
`>=` (lebih besar atau sama dengan), `<=` (lebih kecil atau sama dengan),
`&&` (dan), `||` (atau), `!` (tidak).
*/
package main

import "fmt"

func main() {
	isActor := "Tom Hardy"

	// if-else: satu kondisi atau alternatifnya.
	if isActor == "Tom Hardy" { // Jika isActor sama dengan "Tom Hardy".
		fmt.Println("Good")
	} else { // Jika kondisi di atas salah.
		fmt.Println("Wrong Person")
	}

	singer := "Morgan Wallen"
	whatSongs := "Cover me up"

	// if-else if-else: banyak kondisi berurutan.
	if singer == "Morgan Wallen" || whatSongs == "Whiskey Glasses" { // Jika singer "Morgan Wallen" ATAU lagu "Whiskey Glasses".
		fmt.Println("country man")
	} else if singer == "Morgan Wallen" && whatSongs == "Cover me up" { // Jika singer "Morgan Wallen" DAN lagu "Cover me up".
		fmt.Println("country boy")
	} else { // Jika semua kondisi di atas salah.
		fmt.Println("you not a fan boy or fan man")
	}

	// Nested if: if di dalam if.
	if whatSongs == "Cow Girls" { // Kondisi luar.
		if whatSongs == "Last Night" { // Kondisi dalam 1.
			fmt.Println("good song ")
		} else if whatSongs == "Cover me up" { // Kondisi dalam 2.
			fmt.Println("yes")
		}
	} else { // Jika kondisi luar salah.
		fmt.Println("you dont like country music")
	}
}
