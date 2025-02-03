package main

import "fmt"

/*
Perbedaan Slice dan Array:

1. **Array**:
   - Memiliki jumlah elemen tetap yang harus didefinisikan saat deklarasi.
   - Ukurannya tidak dapat diubah setelah didefinisikan.
   - Contoh deklarasi array:
     var arr = [3]int{1, 2, 3}

2. **Slice**:
   - Tidak memerlukan jumlah elemen yang didefinisikan saat deklarasi.
   - Ukurannya dinamis, dapat bertambah atau berkurang sesuai kebutuhan.
   - Slice adalah referensi ke array di belakang layar (backing array).
   - Contoh deklarasi slice:
     var slc = []int{1, 2, 3}
*/

func main() {
	// Membuat slice integer
	x := []int{1, 2, 3, 4, 5}

	// Membuat slice string
	a := []string{"Tom", "Hardy"}

	// Menampilkan slice awal
	fmt.Println("Slice x:", x)
	fmt.Println("Slice a:", a)

	// Mengubah elemen dalam slice
	x[2] = 10
	a[1] = "Holland" // Menghindari string kosong, bisa diisi string lain

	fmt.Println("Setelah perubahan x[2]:", x)
	fmt.Println("Setelah perubahan a:", a)

	// Menambahkan elemen ke slice menggunakan append
	x = append(x, 6, 7) // Tambahkan elemen baru
	a = append(a, "Hardy")

	fmt.Println("Setelah append x:", x)
	fmt.Println("Setelah append a:", a)

	// Memotong slice (mengubah panjangnya)
	x = x[:4] // Mengambil hanya 4 elemen pertama
	fmt.Println("Slice x setelah dipotong:", x)

	// Copy slice
	originalSlice := []int{100, 200, 300}          // Slice sumber
	copiedSlice := make([]int, len(originalSlice)) // Buat slice dengan panjang yang sama
	copy(copiedSlice, originalSlice)               // Salin data dari originalSlice ke copiedSlice

	fmt.Println("Original slice:", originalSlice)
	fmt.Println("Copied slice:", copiedSlice)
}
