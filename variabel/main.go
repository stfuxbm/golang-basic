package main

import "fmt"

// Deklarasi variabel global secara grup (untuk digunakan di seluruh paket)
/*
Jika variabel tersebut perlu diakses oleh beberapa fungsi dalam program, variabel global lebih tepat digunakan.
*/
var (
	fullName       = "Jhon Due" // Tipe data string
	ages           = 29         // Tipe data int
	isHandsome     = true       // Tipe data bool
	height     int = 171        // Tipe data int dengan nilai eksplisit
)

func main() {
	// Deklarasi variabel tanpa tipe data eksplisit (type inference)
	var stringData = "ini string data" // Tipe otomatis: string
	var booleanData = true             // Tipe otomatis: bool
	var intData = 7                    // Tipe otomatis: int

	// Deklarasi variabel dengan tipe data eksplisit
	var tipeString string = "dengan tipe data" // Tipe eksplisit string
	var tipeBoolean bool = false               // Tipe eksplisit bool
	var tipeInteger int = 1                    // Tipe eksplisit int

	// Deklarasi menggunakan shorthand assignment (tipe otomatis)
	tipeShort := 5 // Tipe otomatis ditentukan oleh Go, bisa int, float, string, dll

	// Deklarasi multi-variabel dengan tipe otomatis
	var x, y, z = "lets go", 1, true // x: string, y: int, z: bool bisa juga dengan shorthand x, y, z := 1, 2, 5

	// Deklarasi variabel tanpa nilai (default value diberikan oleh Go)
	var a string // Tipe string, nilai default adalah "" (kosong)
	var b int    // Tipe int, nilai default adalah 0
	var c bool   // Tipe bool, nilai default adalah false

	// Menampilkan hasil deklarasi variabel
	fmt.Println("Deklarasi tanpa tipe data:", stringData, booleanData, intData)
	fmt.Println("Deklarasi dengan tipe data:", tipeString, tipeBoolean, tipeInteger)
	fmt.Println("Deklarasi dengan shorthand:", tipeShort)
	fmt.Println("Deklarasi dengan multi variabel:", x, y, z)
	fmt.Println("Deklarasi dengan grup global:", fullName, ages, isHandsome, height)
	fmt.Println("Deklarasi tanpa nilai:", a, b, c)
}
