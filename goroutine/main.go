/*
Goroutine: fungsi yang berjalan secara konkuren (bersamaan).
Digunakan untuk menjalankan tugas paralel tanpa memblokir program utama,
ideal untuk I/O, tugas latar belakang, dan meningkatkan responsivitas.
Gunakan kata kunci 'go' untuk menjalankannya.
*/
package main

import (
	"fmt"
	"time"
)

// Cetak pesan berulang kali.
func LoopingPrint(number int, message string) {
	for i := 0; i < number; i++ {
		fmt.Println(i+1, message) // Cetak nomor dan pesan.
	}
}

// Struktur untuk data person.
type Person struct {
	ID   int
	Name string
}

// Method untuk mengembalikan nama (duplikat).
func (p Person) SayMyName() string {
	return p.Name + " dari " + p.Name
}

// Method untuk mencetak ID dan nama.
func (p Person) CallMe() {
	fmt.Println(p.ID, p.Name)
}

func main() {
	// Buat objek Person.
	person := Person{ID: 123, Name: "Jhon Pardy"}

	// Jalankan LoopingPrint sebagai goroutine.
	go LoopingPrint(5, "Yes")

	// Jalankan CallMe sebagai goroutine.
	go person.CallMe()

	// Jalankan SayMyName sebagai goroutine (cetak hasilnya).
	go func() {
		fmt.Println(person.SayMyName())
	}()

	// Tunggu sebentar agar goroutine selesai.
	time.Sleep(time.Second)

	// Selesai.
	fmt.Println("Done")
}
