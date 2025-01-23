package main

import "fmt"

/*
Map adalah tipe data berbentuk key-value pair, di mana setiap value memiliki key unik.
Key digunakan untuk mengakses, mengubah, atau menghapus value yang terkait dengannya.
*/

func main() {
	/*
		Inisialisasi Map Kosong dengan Deklarasi Langsung
		Format: map[tipe_key]tipe_value
	*/
	x := map[string]int{}                    // Inisialisasi map kosong dengan key bertipe string dan value bertipe int
	x["you"] = 20                            // Menambahkan elemen dengan key "you" dan value 20
	x["me"] = 10                             // Menambahkan elemen dengan key "me" dan value 10
	fmt.Println("Map dengan short hand:", x) // Output map yang sudah diisi

	/*
		Inisialisasi Map dengan Nilai Awal
		Deklarasi menggunakan literal untuk langsung memberikan nilai awal.
	*/
	var y = map[string]int{
		"age":   50,  // Key "age" dengan value 50
		"power": 100, // Key "power" dengan value 100
		"level": 90,  // Key "level" dengan value 90
	}
	fmt.Println("Map dengan nilai awal:", y) // Output map dengan nilai awal

	// Iterasi Map menggunakan for-range
	fmt.Println("Iterasi map x:")
	for key, val := range x { // Iterasi melalui map x
		fmt.Printf("%s\t: %d\n", key, val) // Menampilkan key dan value dalam map x
	}

	// Menghapus Elemen dari Map
	fmt.Println("Map sebelum penghapusan:", y)             // Menampilkan map sebelum elemen dihapus
	delete(y, "level")                                     // Fungsi delete menghapus elemen berdasarkan key
	fmt.Println("Map setelah penghapusan key 'level':", y) // Menampilkan map setelah elemen dengan key "level" dihapus

	/*
		Mengakses Item dengan Key Tertentu
		Fungsi ini mengembalikan dua nilai:
		1. Value yang terkait dengan key.
		2. Boolean (true/false) yang menunjukkan apakah key tersebut ada.
	*/
	var j = map[string]bool{
		"isGood":   true,  // Key "isGood" dengan value true
		"isMedium": false, // Key "isMedium" dengan value false
	}
	value, isExist := j["isGood"] // Mencoba mengakses key "isGood" dan memeriksa apakah ada
	if isExist {
		fmt.Println("Value untuk key 'isGood':", value) // Jika key ada, tampilkan value-nya
	} else {
		fmt.Println("Key 'isGood' tidak ditemukan") // Jika key tidak ada, beri pesan
	}

	/*
		Kombinasi Map dan Slice
		Map dapat digunakan sebagai elemen slice untuk menyimpan kumpulan data dengan struktur key-value.
	*/
	brands := []map[string]string{
		{"name": "Lenovo", "country": "China"}, // Map dengan key "name" dan "country"
		{"name": "Asus", "country": "Taiwan"},  // Map dengan key "name" dan "country"
		{"name": "Apple", "country": "USA"},    // Map dengan key "name" dan "country"
	}

	// Menampilkan hanya range dalam map (index array)
	fmt.Println("Menampilkan hanya range dalam map")
	for brand := range brands { // Iterasi melalui slice brands
		fmt.Println(brand) // Output index elemen dalam slice, bukan key-value dari map
	}

	// Menampilkan key dan value dalam map
	fmt.Println("Menampilkan key dan value")
	for _, brand := range brands { // Iterasi untuk mendapatkan elemen map dari slice
		fmt.Printf("Nama: %s, Negara: %s\n", brand["name"], brand["country"]) // Menampilkan key dan value dalam map
	}
}
