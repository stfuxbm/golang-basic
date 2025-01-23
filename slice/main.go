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
	// Contoh deklarasi dan penggunaan array
	var x = []string{
		"hai",
		"halo",
		"yes",
		"no",
	}
	// Menggunakan slicing pada array
	fmt.Println("Array slicing [0:3]:", x[0:3])

	// Contoh deklarasi dan penggunaan slice
	var z = []int{
		1, 2, 3, 4,
	}
	// Menampilkan elemen slice
	fmt.Println("Elemen pertama slice:", z[0])
	fmt.Println("Elemen kedua slice:", z[1])
	// Slicing pada slice
	fmt.Println("Slice slicing [0:1]:", z[0:1])
	fmt.Println("Seluruh elemen slice:", z[:])
	// Panjang slice
	fmt.Println("Panjang slice:", len(z))
}
