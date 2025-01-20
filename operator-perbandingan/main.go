package main

/*
Operator perbandingan digunakan untuk membandingkan 2 nilai.
Hasil dari perbandingan berupa nilai boolean: `true` atau `false`.
- Jika hasil perbandingan benar, maka nilainya `true`.
- Jika hasil perbandingan salah, maka nilainya `false`.
*/

func main() {

	var a int = 10
	var b int = 10

	println("Sama dengan:", a == b)                  // true
	println("Tidak sama dengan:", a != b)            // false
	println("Lebih kecil dari:", a < b)              // false (jika nilainya sama, hasilnya `false`)
	println("Lebih kecil atau sama dengan:", a <= b) // true (jika nilainya sama, hasilnya `true` karena ada operator sama dengan)
	println("Lebih besar dari:", a > b)              // false
	println("Lebih besar atau sama dengan:", a >= b) // true
}
