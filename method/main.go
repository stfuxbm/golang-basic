/*
Struct: kumpulan field (variabel) dengan tipe data berbeda dalam satu unit.
Digunakan untuk merepresentasikan data yang memiliki beberapa atribut terkait.
Mirip seperti objek sederhana atau record data.
*/
package main

import "fmt"

// Struktur untuk data Student.
type Student struct {
	Name  map[string]string // Nama (first dan last).
	Age   int               // Umur.
	Grade []int             // Daftar nilai.
}

// Method untuk menampilkan data Student.
func (s Student) ShowStudents() {
	fmt.Printf("Name: %s %s | Age: %d | Grades: %v\n",
		s.Name["firstName"], s.Name["lastName"], s.Age, s.Grade)
}

// Method untuk mengubah umur Student (menggunakan pointer).
func (s *Student) UpdateAge(newAge int) {
	s.Age = newAge // Mengubah nilai Age.
}

func main() {
	// Membuat slice (mirip list dinamis) dari Student.
	students := []Student{
		{Name: map[string]string{"firstName": "Tom", "lastName": "Hardy"}, Age: 34, Grade: []int{50, 60, 34}},
		{Name: map[string]string{"firstName": "Jhon", "lastName": "Wayne"}, Age: 30, Grade: []int{70, 85, 90}},
	}

	// Tampilkan data sebelum update.
	fmt.Println("\n===== SEBELUM UPDATE =====")
	for _, student := range students {
		student.ShowStudents()
	}

	// Ubah umur semua student jadi 34.
	for i := range students {
		students[i].UpdateAge(34)
	}

	// Ubah umur student pertama jadi 12.
	students[0].UpdateAge(12)

	// Tampilkan data setelah update.
	fmt.Println("\n===== SETELAH UPDATE =====")
	for _, student := range students {
		student.ShowStudents()
	}
}
