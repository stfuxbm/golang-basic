package main

import (
	"fmt"
	"reflect" // Package ini digunakan untuk memeriksa tipe data dari sebuah variabel
)

var x = 1 // Deklarasi variabel global dengan nilai 1

// Deklarasi variabel secara grup, bisa di luar fungsi dan bisa dideklarasikan dengan tipe data tertentu
var (
	fullName       = "Braiyen Massora" // Tipe data string
	ages           = 29                // Tipe data int
	isHandsome     = true              // Tipe data bool
	height     int = 171               // Tipe data int, nilai default bisa langsung ditetapkan
)

// Variabel lokal hanya dapat dideklarasikan menggunakan format ini di dalam fungsi main
// y := 2 hanya bisa digunakan di dalam fungsi main
func main() {
	// DEKLARASI VARIABEL DENGAN NILAI
	fmt.Println("// DEKLARASI VARIABEL DENGAN NILAI")
	// Penamaan variabel di Go mayoritas menggunakan camelCase
	var name = "Braiyen"                   // Tipe data string, Go akan mendeteksi tipe secara otomatis
	age := 17                              // Tipe data int, deklarasi dengan cara ini hanya berlaku di dalam fungsi
	var locationAddress string = "Jakarta" // Tipe data string, hanya menyimpan teks

	// Output untuk menunjukkan nilai variabel
	fmt.Println("Hello my name is: ", name)
	fmt.Println("My age is: ", age)
	fmt.Println("You can find me at: ", locationAddress)
	fmt.Println("Nilai variabel x: ", x)

	// DEKLARASI VARIABEL TANPA NILAI
	fmt.Println("// DEKLARASI VARIABEL TANPA NILAI")
	var a string // Tipe data string tanpa nilai, secara default akan bernilai kosong ""
	var b int    // Tipe data int tanpa nilai, secara default akan bernilai 0
	var c bool   // Tipe data bool tanpa nilai, secara default akan bernilai false

	// Output untuk nilai default dari variabel tanpa nilai
	fmt.Println("Ini akan menghasilkan default string kosong yaitu: ", a)
	fmt.Println("Ini akan menghasilkan default 0 yaitu: ", b)
	fmt.Println("Ini akan menghasilkan default false yaitu: ", c)

	// DEKLARASI MULTIPLE VARIABEL
	fmt.Println("// DEKLARASI MULTIPLE VARIABEL")
	// Deklarasi beberapa variabel dalam satu baris, tipe data harus sudah diketahui
	var d, e, f int = 4, 5, 6
	var cityFrom, cityTo, datePicker string = "Toraja", "Makassar", "12 Januari"

	// Output untuk beberapa variabel yang dideklarasikan sekaligus
	fmt.Println(d, e, f)
	fmt.Println("I think you're from ", cityFrom, "and after meeting you go to ", cityTo, "on ", datePicker)

	// Deklarasi variabel dengan tipe data dinamis, Go akan mendeteksi tipe datanya secara otomatis
	var myHobby, ranking, isWinner = "idk", 1, true

	// Output untuk memeriksa tipe data variabel
	fmt.Println("Tipe data dari myHobby adalah: ", reflect.TypeOf(myHobby), "yaitu ", myHobby)
	fmt.Println("Tipe data dari ranking adalah: ", reflect.TypeOf(ranking), "yaitu ", ranking)
	fmt.Println("Tipe data dari isWinner adalah: ", reflect.TypeOf(isWinner), "yaitu ", isWinner)

	// DEKLARASI MULTIPLE VARIABEL DALAM BLOK
	fmt.Println("// DEKLARASI MULTIPLE VARIABEL DALAM BLOK")
	// Output beberapa variabel dalam satu baris
	fmt.Println("Hi, it's me", fullName, "age", ages, "if you see", isHandsome, "that means I am handsome", "and my height is", height)
}
