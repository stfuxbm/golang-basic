/*
Web Server Sederhana: membuat aplikasi web yang melayani konten melalui HTTP.
Digunakan untuk membangun aplikasi web, API, atau layanan jaringan lainnya.
Paket 'net/http' menyediakan fungsi dasar untuk membuat server web.
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler untuk halaman "/about".
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About") // Kirim teks "About" sebagai response.
}

// Handler untuk halaman "/contact".
func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contact") // Kirim teks "Contact" sebagai response.
}

// Struktur data Person.
type Person struct {
	Id      string
	Name    string
	Country string
}

// Method untuk menampilkan informasi Person.
func (p Person) ShowPerson() string {
	return fmt.Sprintf("ID: %s, Name: %s, Country: %s", p.Id, p.Name, p.Country)
}

func main() {
	// Membuat instance Person.
	person := Person{Id: "123", Name: "Jhon", Country: "Indonesia"}

	// Handler untuk halaman "/".
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Ini Halaman Home") // Kirim teks "Ini Halaman Home".
	})

	// Hubungkan path "/about" dengan handler fungsi 'about'.
	http.HandleFunc("/about", about)

	// Hubungkan path "/contact" dengan handler fungsi 'contact'.
	http.HandleFunc("/contact", contact)

	// Handler untuk halaman "/users".
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, person.ShowPerson()) // Kirim informasi person.
	})

	// Start web server di port 8080.
	fmt.Println("Starting server at http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err) // Catat error jika server gagal start.
	}
}
