/*
Go channel adalah mekanisme komunikasi antar goroutine yang aman dan terstruktur.
Channel memungkinkan satu goroutine mengirim nilai ke goroutine lain.
Mereka adalah pipa yang menghubungkan goroutine, memungkinkan pertukaran data.

Kapan Go channel digunakan?
- Untuk mengirim data antar goroutine secara aman (menghindari race condition).
- Untuk melakukan sinkronisasi antar goroutine, di mana satu goroutine menunggu data dari goroutine lain.
- Dalam pola desain concurrency seperti pipeline dan worker pool.
- Ketika Anda perlu mengkoordinasikan pekerjaan yang dilakukan secara paralel.

Secara singkat, channel digunakan untuk komunikasi dan sinkronisasi antar bagian kode yang berjalan secara bersamaan (goroutine).
*/
package main

import (
	"fmt"
	"time"
)

func sendMessage(ch chan string, msg string) {
	// Mengirim pesan ke channel setelah penundaan singkat.
	time.Sleep(time.Millisecond * 100)
	ch <- msg
}

func receiveMessage(ch chan string) {
	// Menerima pesan dari channel dan mencetaknya.
	msg := <-ch
	fmt.Println("Menerima:", msg)
}

func main() {
	// Membuat channel 'messages' yang dapat mengirim dan menerima nilai bertipe string.
	messages := make(chan string)

	// Menjalankan goroutine untuk mengirim pesan.
	go sendMessage(messages, "Halo dari goroutine 1")
	go sendMessage(messages, "Pesan kedua dari goroutine 2")

	// Menjalankan goroutine untuk menerima pesan.
	go receiveMessage(messages)
	go receiveMessage(messages)

	// Memberikan waktu agar goroutine selesai berjalan sebelum program utama berakhir.
	time.Sleep(time.Second)
	fmt.Println("Selesai")
}
