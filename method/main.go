package main

import "fmt"

// Struct Student dengan Name (map), Age (int), dan Grade (slice int)
type Student struct {
	Name  map[string]string
	Age   int
	Grade []int
}

// Method untuk menampilkan data Student dengan format lebih rapi
func (s Student) ShowStudents() {
	fmt.Printf("Name: %s %s | Age: %d | Total Grades: %d | Grades: %v\n",
		s.Name["firstName"], s.Name["lastName"], s.Age, len(s.Grade), s.Grade)
}

// Method untuk update age menggunakan pointer
func (s *Student) UpdateAge(newAge int) {
	s.Age = newAge
}

func main() {
	// Membuat slice of Student
	students := []Student{
		{
			Name: map[string]string{"firstName": "Tom", "lastName": "Hardy"},
			Age:  34,
			Grade: []int{
				50, 60, 34,
			},
		},
		{
			Name: map[string]string{"firstName": "Jhon", "lastName": "Wayne"},
			Age:  30,
			Grade: []int{
				70, 85, 90,
			},
		},
	}

	// Menampilkan semua data sebelum update
	fmt.Println("\n===== SEBELUM UPDATE DATA =====")
	for _, student := range students {
		student.ShowStudents()
	}

	// Mengubah semua umur menjadi 34
	for i := range students {
		students[i].UpdateAge(34)
	}

	// Mengubah umur hanya untuk index pertama menjadi 12
	students[0].UpdateAge(12)

	// Menampilkan semua data setelah update
	fmt.Println("\n===== SETELAH UPDATE DATA =====")
	for _, student := range students {
		student.ShowStudents()
	}
}
