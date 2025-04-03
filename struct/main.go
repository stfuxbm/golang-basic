/*
Struct: tipe data komposit untuk mengelompokkan nilai dengan tipe berbeda.
Digunakan untuk merepresentasikan entitas dengan beberapa atribut.
Bisa berisi struct lain (nested), slice, atau map.
Method dapat diasosiasikan dengan struct untuk perilaku objek.
*/
package main

import (
	"fmt"
)

// Struct dengan struct di dalamnya (Nested Struct).
type Address struct {
	Country string
	City    string
}
type Person struct {
	ID          string
	Name        string
	Email       string
	PhoneNumber int
	Address     Address
}

// Struct dengan array/slice.
type Student struct {
	Nama  string
	Grade []int
}

// Struct dengan map (tipe data dinamis).
type Teachers struct {
	Name    string
	Country map[string]string
}

// Method untuk Nested Struct.
func (j Person) ExampleNestedStruct() {
	fmt.Println(j.Name)
	fmt.Println(j.Address.Country)
}

// Method untuk struct Person.
func (y Person) ExampleStructFunction() {
	fmt.Println("Name:", y.Name, ", Email:", y.Email)
}

// Method untuk struct Student.
func (s Student) ExampleStructArraySlice() {
	fmt.Println("Name:", s.Nama)
	fmt.Println("Grades:", s.Grade)
	fmt.Println("Total Subjects:", len(s.Grade))
}

// Method untuk struct Teachers.
func (t Teachers) ExampleStructDataDynamis() {
	fmt.Println("Name:", t.Name)
	fmt.Println("Country:", t.Country)
	fmt.Println("Akses data map:", t.Country["Negara Asal"])
}

func main() {
	// Slice dari struct Person.
	x := []Person{
		{ID: "123", Name: "Tom Hardy", Email: "tom@hardy.com", PhoneNumber: 123},
		{ID: "124", Name: "Keanu Reeves", Email: "keanu@reeves.com", PhoneNumber: 456},
	}

	// Instance struct Student.
	a := Student{Nama: "James", Grade: []int{60, 80, 70}}

	// Instance struct Teachers dengan map.
	t := Teachers{Name: "Peter", Country: map[string]string{"Negara Asal": "Swedia", "Negara Domisili": "Indonesia"}}

	// Instance struct Person dengan nested struct Address.
	j := Person{ID: "125", Name: "Chris Hemsworth", Email: "chris@thor.com", PhoneNumber: 789,
		Address: Address{Country: "Australia", City: "Melbourne"}}

	fmt.Println("\n-----Nested Struct Method-----")
	j.ExampleNestedStruct()

	fmt.Println("\n-----Struct dengan Map Method-----")
	t.ExampleStructDataDynamis()

	fmt.Println("\n-----Struct dengan Slice Method-----")
	a.ExampleStructArraySlice()

	fmt.Println("\n-----Method untuk Setiap Struct Person-----")
	for _, person := range x {
		person.ExampleStructFunction()
	}

	fmt.Println("\n-----Menampilkan dengan If Else-----")
	for _, users := range x {
		if users.Name == "Tom Hardy" {
			fmt.Println("Nama ditemukan:", users.Name)
		} else if users.Name == "Keanu Reeves" {
			fmt.Println("Nama ditemukan:", users.Name)
		} else {
			fmt.Println("Nama tidak ditemukan, yang ditemukan:", users.Name)
		}
	}
}
