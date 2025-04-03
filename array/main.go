/*
Goroutine: fungsi yang berjalan secara konkuren (bersamaan).
Digunakan untuk menjalankan tugas secara paralel tanpa memblokir program utama.
Ideal untuk operasi I/O, tugas latar belakang, dan meningkatkan responsivitas aplikasi.
*/
package main

import (
	"fmt"
	"time"
)

func LoopingPrint(number int, message string) {
	// Cetak pesan berulang kali.
	for i := 0; i < number; i++ {
		fmt.Println(i+1, message)
	}
}

type Person struct {
	ID   int
	Name string
}

func (p Person) SayMyName() string {
	// Mengembalikan nama (sebenarnya duplikasi nama).
	return p.Name + " dari " + p.Name
}

func (p Person) CallMe() {
	// Mencetak ID dan nama person.
	fmt.Println(p.ID, p.Name)
}

func main() {
	// Membuat instance struct Person.
	person := Person{
		ID:   123,
		Name: "Jhon Pardy",
	}

	// Jalankan LoopingPrint sebagai goroutine.
	go LoopingPrint(5, "Yes")

	// Jalankan method CallMe sebagai goroutine.
	go person.CallMe()

	// Jalankan method SayMyName sebagai goroutine (hasil tidak digunakan).
	go person.SayMyName()

	// Tunggu sebentar agar goroutine punya waktu berjalan.
	time.Sleep(time.Second)

	// Cetak "Done" setelah jeda.
	fmt.Println("Done")
}
