package main

import (
	"fmt"
)

// ================= INTERFACE =================

// Interface HasName: Digunakan untuk mendapatkan informasi seseorang
// Struct yang mengimplementasikan interface ini harus memiliki kedua method ini.
type HasName interface {
	GetPeopleInformation() string
	GetPeopleIdentify() string
}

// Interface Role: Digunakan untuk mendapatkan peran seseorang
type Role interface {
	GetRole() string
}

// Interface EmbedInterface: Menggabungkan beberapa interface
// Struct yang mengimplementasikan interface ini harus memiliki semua method dari HasName dan Role
type EmbedInterface interface {
	ID() string
	HasName // Meng-embed interface HasName
	Role    // Meng-embed interface Role
}

// ================= STRUCT PERSON =================

// Struct Person mengimplementasikan interface EmbedInterface
type Person struct {
	IDValue string // ID dari orang tersebut
	Name    string // Nama orang
	Country string // Negara asal
	Sex     bool   // Jenis kelamin (true = Male, false = Female)
	Role    string // Peran orang tersebut
}

// ================= IMPLEMENTASI METHOD =================

// Implementasi method GetPeopleInformation() dari interface HasName
func (p Person) GetPeopleInformation() string {
	return p.Name + " dari " + p.Country
}

// Implementasi method GetPeopleIdentify() dari interface HasName
func (p Person) GetPeopleIdentify() string {
	if p.Sex {
		return "Male"
	}
	return "Female"
}

// Implementasi method GetRole() dari interface Role
func (p Person) GetRole() string {
	return p.Role
}

// Implementasi method ID() dari interface EmbedInterface
func (p Person) ID() string {
	return p.IDValue
}

// ================= FUNGSI MENGGUNAKAN INTERFACE =================

// Fungsi SayHello menerima parameter dengan interface EmbedInterface
func SayHello(e EmbedInterface) {
	fmt.Println("Halo", e.GetPeopleInformation())
	fmt.Println("Identify as:", e.GetPeopleIdentify())
	fmt.Println("Role:", e.GetRole())
	fmt.Println("ID:", e.ID())
}

// ================= INTERFACE KOSONG =================

// Interface kosong dapat menerima nilai dengan tipe data apa pun.
type ExampleInterfaceEmpty interface{} // Bisa juga menggunakan `any` di Go 1.18+

// Fungsi untuk menampilkan nilai dari interface kosong
func CallExampleInterfaceEmpty(e ExampleInterfaceEmpty) {
	fmt.Println("Interface Empty Value:", e)
}

// ================= FUNGSI UTAMA (MAIN) =================

func main() {
	// Membuat objek dari struct Person
	person := Person{
		IDValue: "12345",
		Name:    "Tom Hardy",
		Country: "England",
		Sex:     true,
		Role:    "Actor",
	}

	// Memanggil fungsi SayHello() dengan objek person
	SayHello(person)

	fmt.Println("\n=== Contoh Penggunaan Interface Kosong ===")

	// Memanggil fungsi CallExampleInterfaceEmpty() dengan berbagai tipe data
	CallExampleInterfaceEmpty(true)                                       // Boolean
	CallExampleInterfaceEmpty("John")                                     // String
	CallExampleInterfaceEmpty(123)                                        // Integer
	CallExampleInterfaceEmpty([]string{"Actor", "Musician", "President"}) // Slice

}
