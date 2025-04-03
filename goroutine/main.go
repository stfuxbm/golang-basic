package main

/*
Goroutine di Go sangat berguna dalam berbagai situasi yang membutuhkan concurrency (eksekusi bersamaan) atau paralelisme (pekerjaan yang dibagi di beberapa core atau thread),
tanpa membuat aplikasi menjadi kompleks atau memakan banyak sumber daya.
*/

import (
	"fmt"
	"time"
)

// Fungsi LoopingPrint mencetak pesan sebanyak 'number' kali
func LoopingPrint(number int, message string) {
	for i := 0; i < number; i++ {
		fmt.Println(i+1, message) // Mencetak nomor urut dan pesan
	}
}

// Struktur Person yang memiliki dua field: ID dan Name
type Person struct {
	ID   int
	Name string
}

// Metode SayMyName mengembalikan string yang berisi nama orang
func (p Person) SayMyName() string {
	return p.Name + " dari " + p.Name
}

// Metode CallMe mencetak ID dan nama orang
func (p Person) CallMe() {
	fmt.Println(p.ID, p.Name)
}

func main() {
	// Membuat objek Person dengan ID 123 dan nama "Jhon Pardy"
	person := Person{
		ID:   123,
		Name: "Jhon Pardy",
	}

	// Memulai goroutine untuk mencetak pesan "Yes" sebanyak 5 kali
	go LoopingPrint(5, "Yes")

	// Memulai goroutine untuk mencetak ID dan nama orang
	go person.CallMe()

	// Memulai goroutine untuk mengembalikan string nama orang
	go func() {
		fmt.Println(person.SayMyName())
	}()

	// Memberi waktu untuk goroutine selesai sebelum program berakhir
	time.Sleep(time.Second)

	// Mencetak pesan "Done" setelah semua goroutine selesai
	fmt.Println("Done")
}
