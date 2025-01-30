package main

import "fmt"

// Fungsi dengan dua nilai kembali (multiple return values)
// Fungsi ini menerima integer dan string, lalu mengembalikan hasil perhitungan dan string yang dimodifikasi.
func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x       // Menghitung hasil penjumlahan x + x
	txt1 = y + " World!" // Menambahkan kata "World!" ke dalam string y
	return               // Mengembalikan hasil perhitungan dan string yang telah dimodifikasi
}

// Fungsi untuk menghitung usia berdasarkan tahun kelahiran
// Fungsi ini menerima tahun lahir (x), nama (y), dan status kepemilikan (owner)
// Mengembalikan umur (int), nama (string), dan status owner (bool)
func showYourID(x int, y string, owner bool) (int, string, bool) {
	age := 2025 - x           // Menghitung usia berdasarkan tahun kelahiran
	name := y                 // Menyimpan nama yang diberikan
	isOwner := owner          // Menyimpan status kepemilikan
	return age, name, isOwner // Mengembalikan tiga nilai: usia, nama, dan status owner
}

func main() {
	// Memanggil myFunction dan menampilkan hasilnya
	fmt.Println(myFunction(5, "Hello")) // Output: (10, "Hello World!")

	// Memanggil showYourID dengan tahun kelahiran, nama, dan status owner
	age, name, isOwner := showYourID(1995, "Jhon", true)
	fmt.Println("Age:", age)       // Output: Age: 30
	fmt.Println("Name:", name)     // Output: Name: Jhon
	fmt.Println("Owner:", isOwner) // Output: Owner: true

	// Contoh tambahan fungsi sederhana
	fmt.Println(square(4))            // Output: 16
	fmt.Println(greet("Alice", "Hi")) // Output: Hi, Alice!
}

// Contoh 1: Fungsi sederhana untuk menghitung kuadrat dari suatu angka
func square(num int) int {
	return num * num // Mengembalikan hasil kuadrat dari angka
}

// Contoh 2: Fungsi untuk membuat salam dari nama dan pesan
func greet(name string, greeting string) string {
	return greeting + ", " + name + "!" // Mengembalikan string salam yang sudah digabungkan
}
