/*
URL Parsing: memecah string URL menjadi komponen-komponennya.
Digunakan untuk menganalisis dan memanipulasi bagian-bagian URL,
seperti skema, host, path, dan query parameters.
Paket 'net/url' menyediakan fungsi untuk ini.
*/
package main

import (
	"fmt"
	"net/url"
)

func main() {
	// String URL yang akan di-parse.
	urlString := "https://m.borlindo.com//order/?departureCity=MKS&destinationCity=TRJ&departureDate=2025-04-30&isHorizontal=true&version=2"

	// Parse URL (kembalikan URL ter-parse dan error).
	u, e := url.Parse(urlString)

	// Cek jika ada error saat parsing.
	if e != nil {
		fmt.Println(e.Error()) // Cetak error.
		return
	}

	fmt.Println("URL:", urlString)
	fmt.Println("Host:", u.Host) // Domain.
	fmt.Println("Path:", u.Path) // Bagian setelah domain.

	// Ambil query parameter.
	departureCity := u.Query()["departureCity"]
	destinationCity := u.Query()["destinationCity"]

	fmt.Println("departureCity:", departureCity)
	fmt.Println("destinationCity:", destinationCity)
}
