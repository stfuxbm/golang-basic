package main

import (
	"fmt"
)

func main() {
	// Membuat channel yang dapat mengirim dan menerima nilai bertipe string
	message := make(chan string)

	// Fungsi anonim untuk mengirim pesan melalui channel
	var SayMyName = func(who string) {
		data := fmt.Sprintf("Hai %s", who)
		message <- data // Mengirim data ke channel
	}

	// Menjalankan beberapa goroutine
	go SayMyName("Jhon")
	go SayMyName("Tom")
	go SayMyName("Hardy")

	// Menerima pesan dari channel sebanyak jumlah goroutine
	// Hasil bisa saja tidak berurutan karena goroutine berjalan secara paralel
	for i := 0; i < 3; i++ {
		fmt.Println(<-message)
	}

	// Menutup channel karena sudah tidak digunakan lagi
	close(message)
}
