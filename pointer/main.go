package main

import "fmt"

// Struct Person
type Person struct {
	Name    string
	Age     int
	Country string
}

// Method untuk mengubah umur
func (p *Person) ChangeAge(newAge int) {
	p.Age = newAge
}

// Method untuk mengubah nama
func (p *Person) ChangeName(newName string) {
	p.Name = newName
}

// Method untuk mengubah negara
func (p *Person) ChangeCountry(newCountry string) {
	p.Country = newCountry
}

func main() {
	x := Person{
		Name:    "Tom Hardy",
		Age:     34,
		Country: "England",
	}

	fmt.Println("Sebelum:", x)

	// Mengubah nilai dengan pointer method
	x.ChangeAge(40)
	x.ChangeName("Chris Hemsworth")
	x.ChangeCountry("Australia")
	fmt.Println("Sesudah:", x)

	// Cek alamat memori
	y := &x
	fmt.Println("\nPointer y menunjuk ke x:", y)
	fmt.Println("Nilai yang ditunjuk y:", y.Name, y.Age, y.Country)

	// Menampilkan alamat memori dengan &
	fmt.Println("\nAlamat Memori Struct x:", &x)
	fmt.Println("Alamat Memori Name:", &x.Name)
	fmt.Println("Alamat Memori Age:", &x.Age)
	fmt.Println("Alamat Memori Country:", &x.Country)

	// Menampilkan nilai struct melalui pointer y
	fmt.Println("\nNilai Struct melalui Pointer y:", *y)
}
