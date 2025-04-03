package main

import (
	"fmt"
)

// Define struct Person

// Nested Struct
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

// Define struct dengan array/slice
type Student struct {
	Nama  string
	Grade []int
}

// Define struct untuk tipe data dinamis
type Teachers struct {
	Name    string
	Country map[string]string
}

// Method  Struct di dalam Struct (Nested Struct)
func (j Person) ExampleNestedStruct() {
	fmt.Println(j.Name)
	fmt.Println(j.Address.Country)
}

// Method untuk struct Person
func (y Person) ExampleStructFunction() {
	fmt.Println("Name:", y.Name, ", Email:", y.Email)
}

// Method untuk struct Student
func (s Student) ExampleStructArraySlice() {
	fmt.Println("Name:", s.Nama)
	fmt.Println("Grades:", s.Grade)
	fmt.Println("Total Subjects:", len(s.Grade)) // Menampilkan jumlah nilai
}

// Method untuk struct Teachers
func (t Teachers) ExampleStructDataDynamis() {
	fmt.Println("Name:", t.Name)
	fmt.Println("Country:", t.Country) // Perbaiki typo
	fmt.Println("Akses data dalam map:", t.Country["Negara Asal"])
}

func main() {
	// Slice dengan lebih dari satu data
	x := []Person{
		{ID: "123", Name: "Tom Hardy", Email: "tom@hardy.com", PhoneNumber: 123},
		{ID: "124", Name: "Keanu Reeves", Email: "keanu@reeves.com", PhoneNumber: 456},
	}

	a := Student{
		Nama:  "James",
		Grade: []int{60, 80, 70},
	}

	t := Teachers{
		Name: "Peter",
		Country: map[string]string{
			"Negara Asal":     "Swedia",
			"Negara Domisili": "Indonesia",
		},
	}
	j := Person{
		ID:          "125",
		Name:        "Chris Hemsworth",
		Email:       "chris@thor.com",
		PhoneNumber: 789,
		Address: Address{
			Country: "Australia",
			City:    "Melbourne",
		},
	}

	// panggil menthod  Struct di dalam Struct (Nested Struct)
	fmt.Println("\n-----Panggil  Struct di dalam Struct (Nested Struct)")
	j.ExampleNestedStruct()
	// Panggil method struct Teachers
	fmt.Println("\n-----Panggil dengan Data Dinamis-----")
	t.ExampleStructDataDynamis()

	// Panggil method struct Student
	fmt.Println("\n----Panggil Struct Student-----")
	a.ExampleStructArraySlice()

	// Panggil method struct untuk setiap Person
	fmt.Println("\n-----Panggil method struct untuk setiap Person-----")
	for _, person := range x {
		person.ExampleStructFunction()
	}

	// Use case switch
	fmt.Println("\n-----Menampilkan dengan Case Switch-----")
	for _, person := range x {
		switch person.Name {
		case "Tom Hardy":
			fmt.Println("Ditemukan:", person.Name)
		case "Keanu Reeves":
			fmt.Println("Ditemukan:", person.Name)
		default:
			fmt.Println("Nama tidak ditemukan, yang ditemukan:", person.Name)
		}
	}

	// Use case if-else lebih optimal
	fmt.Println("\n-----Menampilkan dengan IF ELSE-----")
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
