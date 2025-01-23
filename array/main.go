package main

import "fmt"

/*
Array adalah kumpulan data bertipe sama yang disimpan dalam satu variabel.
Array memiliki panjang tetap yang harus ditentukan saat deklarasi,
kecuali menggunakan `[...]` yang otomatis menghitung panjang berdasarkan elemen awal.
*/
func main() {

	// Deklarasi array tanpa nilai awal
	var x [2]string
	x[0] = "good"
	x[1] = "poor"

	var z [1]bool
	z[0] = true

	fmt.Println("Elemen kedua dari x:", x[1])
	fmt.Println("Elemen pertama dari z:", z[0])

	// Deklarasi array dengan nilai awal
	var a = [2]string{"yes", "no"}
	fmt.Println("Elemen kedua dari a:", a[1])
	fmt.Println("Jumlah elemen dalam array a:", len(a))

	// Deklarasi array tanpa jumlah elemen (panjang otomatis)
	var m = [...]int{1, 2, 3}
	fmt.Println("Elemen ketiga dari m:", m[2])
	fmt.Println("Jumlah elemen dalam array m:", len(m))

	// Perulangan array dengan for
	var brand = [...]string{
		"Lenovo",
		"Asus",
		"Samsung",
	}
	for i := 0; i < len(brand); i++ {
		fmt.Printf("Elemen ke-%d dari brand: %s\n", i, brand[i])
	}

	// Perulangan array dengan for range
	var power = [...]int{1, 2, 4}
	for i, value := range power { // Menggunakan `value` untuk menghindari kebingungan
		fmt.Printf("Indeks: %d, Nilai: %d\n", i, value)
	}

	// Abaikan indeks menggunakan _,
	for _, value := range power {
		fmt.Printf("Nilai (tanpa indeks): %d\n", value)
	}
}
