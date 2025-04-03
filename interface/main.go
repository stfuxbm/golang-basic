/*
Interface: kumpulan definisi method tanpa implementasi.
Menentukan kontrak perilaku yang harus dipenuhi oleh tipe data lain.
Digunakan untuk mencapai polimorfisme dan fleksibilitas kode.
Tipe yang mengimplementasikan interface harus memiliki semua method yang didefinisikan.
*/
package main

import (
	"fmt"
)

// Interface untuk informasi nama.
type HasName interface {
	GetPeopleInformation() string // Mendapatkan info orang.
	GetPeopleIdentify() string    // Mendapatkan identitas orang.
}

// Interface untuk peran.
type Role interface {
	GetRole() string // Mendapatkan peran.
}

// Interface gabungan dari HasName dan Role.
type EmbedInterface interface {
	ID() string // Mendapatkan ID.
	HasName     // Embed interface HasName.
	Role        // Embed interface Role.
}

// Struktur Person mengimplementasikan EmbedInterface.
type Person struct {
	IDValue string
	Name    string
	Country string
	Sex     bool
	Role    string
}

// Implementasi GetPeopleInformation untuk Person.
func (p Person) GetPeopleInformation() string {
	return p.Name + " dari " + p.Country
}

// Implementasi GetPeopleIdentify untuk Person.
func (p Person) GetPeopleIdentify() string {
	if p.Sex {
		return "Male"
	}
	return "Female"
}

// Implementasi GetRole untuk Person.
func (p Person) GetRole() string {
	return p.Role
}

// Implementasi ID untuk Person.
func (p Person) ID() string {
	return p.IDValue
}

// Fungsi menerima interface EmbedInterface.
func SayHello(e EmbedInterface) {
	fmt.Println("Halo", e.GetPeopleInformation())
	fmt.Println("Identify as:", e.GetPeopleIdentify())
	fmt.Println("Role:", e.GetRole())
	fmt.Println("ID:", e.ID())
}

// Interface kosong, bisa menerima tipe data apa pun (mirip 'any').
type ExampleInterfaceEmpty interface{}

// Fungsi menerima interface kosong.
func CallExampleInterfaceEmpty(e ExampleInterfaceEmpty) {
	fmt.Println("Interface Empty Value:", e)
}

func main() {
	// Membuat objek Person.
	person := Person{IDValue: "12345", Name: "Tom Hardy", Country: "England", Sex: true, Role: "Actor"}

	// Memanggil fungsi dengan interface.
	SayHello(person)

	fmt.Println("\n=== Contoh Interface Kosong ===")
	CallExampleInterfaceEmpty(true)
	CallExampleInterfaceEmpty("John")
	CallExampleInterfaceEmpty(123)
	CallExampleInterfaceEmpty([]string{"Actor", "Musician", "President"})
}
