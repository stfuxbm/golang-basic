/*
Pointer: variabel yang menyimpan alamat memori variabel lain.
Digunakan untuk memanipulasi nilai asli variabel secara langsung,
terutama berguna untuk mengubah struct dalam fungsi/method.
Deklarasikan dengan '*' saat tipe, dapatkan alamat dengan '&'.
*/
package main

import "fmt"

// Struktur Person.
type Person struct {
	Name    string
	Age     int
	Country string
}

// Method mengubah umur (menggunakan pointer).
func (p *Person) ChangeAge(newAge int) {
	p.Age = newAge // Mengubah nilai Age langsung.
}

// Method mengubah nama (menggunakan pointer).
func (p *Person) ChangeName(newName string) {
	p.Name = newName // Mengubah nilai Name langsung.
}

// Method mengubah negara (menggunakan pointer).
func (p *Person) ChangeCountry(newCountry string) {
	p.Country = newCountry // Mengubah nilai Country langsung.
}

func main() {
	// Membuat instance struct Person.
	x := Person{Name: "Tom Hardy", Age: 34, Country: "England"}
	fmt.Println("Sebelum:", x)

	// Memanggil method pointer untuk mengubah nilai.
	x.ChangeAge(40)
	x.ChangeName("Chris Hemsworth")
	x.ChangeCountry("Australia")
	fmt.Println("Sesudah:", x)

	// Mendapatkan pointer ke struct x.
	y := &x
	fmt.Println("\nPointer y ke x:", y)
	fmt.Println("Nilai via y:", y.Name, y.Age, y.Country) // Akses field via pointer.

	// Menampilkan alamat memori.
	fmt.Println("\nAlamat x:", &x)
	fmt.Println("Alamat Name:", &x.Name)
	fmt.Println("Alamat Age:", &x.Age)
	fmt.Println("Alamat Country:", &x.Country)

	// Menampilkan nilai struct melalui dereferensi pointer.
	fmt.Println("\nNilai via *y:", *y)
}
