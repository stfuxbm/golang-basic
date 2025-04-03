/*
Variabel: tempat menyimpan data dengan nama.
Digunakan untuk menyimpan nilai yang bisa diubah selama program berjalan.
Deklarasi lokal di dalam fungsi, global di luar fungsi.
*/
package main

import "fmt"

// Variabel global (di luar fungsi main).
var (
	fullName       = "Jhon Due" // string
	ages           = 29         // int
	isHandsome     = true       // bool
	height     int = 171        // int (eksplisit)
)

func main() {
	// Deklarasi lokal tanpa tipe eksplisit (inferensi).
	var stringData = "ini string" // string
	var booleanData = true        // bool
	var intData = 7               // int

	// Deklarasi lokal dengan tipe eksplisit.
	var tipeString string = "tipe string" // string
	var tipeBoolean bool = false          // bool
	var tipeInteger int = 1               // int

	// Deklarasi singkat lokal (hanya di dalam fungsi).
	tipeShort := 5 // int

	// Deklarasi banyak variabel lokal.
	var x, y, z = "go", 1, true // string, int, bool

	// Deklarasi tanpa nilai (nilai default).
	var a string // "" (string kosong)
	var b int    // 0
	var c bool   // false

	fmt.Println("Tanpa tipe:", stringData, booleanData, intData)
	fmt.Println("Dengan tipe:", tipeString, tipeBoolean, tipeInteger)
	fmt.Println("Shorthand:", tipeShort)
	fmt.Println("Multi var:", x, y, z)
	fmt.Println("Global:", fullName, ages, isHandsome, height)
	fmt.Println("Default:", a, b, c)
}
