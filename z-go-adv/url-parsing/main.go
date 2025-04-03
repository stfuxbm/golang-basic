package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Mendeklarasikan sebuah string URL yang ingin diparse
	var urlString = "https://m.borlindo.com//order/?departureCity=MKS&destinationCity=TRJ&departureDate=2025-04-30&isHorizontal=true&version=2"

	// Parse URL, u adalah hasil URL yang terparse, e adalah error (jika ada)
	var u, e = url.Parse(urlString)

	// Mengecek apakah ada error saat parsing URL
	if e != nil {
		// Jika ada error, cetak pesan error dan keluar dari fungsi
		fmt.Println(e.Error())
		return
	}

	// Jika parsing berhasil, tampilkan hasilnya
	fmt.Println("url string: \n", urlString)

	// Menampilkan host (misalnya domain) dari URL yang telah diparse
	fmt.Println("host: \n", u.Host)

	// Menampilkan path (misalnya bagian setelah domain) dari URL yang telah diparse
	fmt.Println("path: \n", u.Path)

	// Mengambil query parameter departureCity dari URL yang diparse
	// u.Query() mengembalikan map yang berisi key dan value dari query string URL
	x := u.Query()["departureCity"]

	// Mengambil query parameter destinationCity dari URL yang diparse
	y := u.Query()["destinationCity"]

	// Menampilkan nilai dari query parameters yang diambil
	fmt.Println("departureCity: \n", x)
	fmt.Println("destinationCity: \n", y)
}
