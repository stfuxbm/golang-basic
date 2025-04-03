package main

import (
	"fmt"
	"log"
	"net/http"
)

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contact")
}

type Person struct {
	Id      string
	Name    string
	Country string
}

func (p Person) ShowPerson() string {
	return fmt.Sprintf("ID: %s, Name: %s, Country: %s", p.Id, p.Name, p.Country)
}

func main() {
	person := Person{
		Id:      "123",
		Name:    "Jhon",
		Country: "Indonesia",
	}

	// Halaman home
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Ini Halaman Home")
	})

	// Halaman about
	http.HandleFunc("/about", about)

	// Halaman Contact
	http.HandleFunc("/contact", contact)

	// Halaman users
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, person.ShowPerson())
	})

	// Menjalankan server pada port 8080
	fmt.Println("Starting web server at http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
