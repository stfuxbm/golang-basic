package main

import "fmt"

/*
   Map adalah struktur data dalam Golang yang menyimpan pasangan key-value.
   - Key bersifat unik dan digunakan untuk mengakses value terkait.
   - Map sangat cepat dalam operasi pencarian, penambahan, dan penghapusan data.
*/

func main() {
	// Membuat map dengan short-hand declaration
	actor := map[string]string{
		"Name": "Tom Hardy",
		"Jobs": "Actor",
	}

	movie := map[int]string{
		2024: "Venom",
		2015: "Legend",
	}

	fmt.Println(actor)
	fmt.Println("Nama aktor:", actor["Name"]) // Mengakses nilai menggunakan key
	fmt.Println(movie)
	fmt.Println("Film tahun 2015:", movie[2015])

	// Membuat map menggunakan function make
	netWorth := make(map[int]string)
	netWorth[2024] = "$55 million"
	fmt.Println("Kekayaan bersih tahun 2024:", netWorth[2024])

	// Menampilkan jumlah elemen dalam map
	fmt.Println("Jumlah elemen di actor:", len(actor))
	fmt.Println("Jumlah elemen di movie:", len(movie))

	// Membuat map kosong
	emptyMapShortHand := map[string]int{} // Menggunakan short-hand declaration
	emptyMapMake := make(map[string]bool) // Menggunakan function make

	fmt.Println("Map kosong dengan shorthand:", emptyMapShortHand)
	fmt.Println("Map kosong dengan make:", emptyMapMake)

	// Menambahkan nilai ke dalam map
	brand := make(map[string]bool)
	brand["isGoodProduct"] = false
	brand["isGoodToBuy"] = true
	brand["isAwesome"] = true
	fmt.Println("Brand setelah penambahan nilai:", brand)

	// Mengubah nilai dalam map
	brand["isGoodToBuy"] = false
	fmt.Println("Brand setelah perubahan nilai:", brand)

	// Menghapus nilai dari map
	delete(brand, "isGoodProduct")
	fmt.Println("Brand setelah penghapusan nilai:", brand)

	// Map bersifat reference type
	/*
	   Map dalam Go adalah reference type, artinya:
	   - Jika sebuah map diberikan ke variabel lain, kedua variabel akan merujuk ke data yang sama.
	   - Perubahan pada salah satu variabel akan berdampak pada yang lain.
	*/
	fmt.Println("Map sebelum reference assignment:", brand)

	newBrand := brand // Menggunakan referensi yang sama
	newBrand["isAwesome"] = false

	fmt.Println("Map setelah reference assignment:", brand)
	fmt.Println("NewBrand setelah mengubah nilai:", newBrand)
}
