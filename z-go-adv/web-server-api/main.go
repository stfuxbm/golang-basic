// Contoh sederhana Web Service API menggunakan Go yang menghasilkan response berupa JSON.
// Package "net/http" digunakan untuk membuat server HTTP dan menangani request,
// sedangkan "encoding/json" digunakan untuk melakukan encoding data Go menjadi format JSON.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// DigitalMusic merepresentasikan struktur data musik digital.
// Tag `json:"nama_field"` digunakan untuk menentukan nama field pada JSON output.
type DigitalMusic struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Years int    `json:"Years"`
}

// artist menyimpan data digital music.
var artist = []DigitalMusic{
	{"1", "Spotify", 2019},
	{"2", "Apple Music", 2010},
	{"3", "Youtube Music", 2008},
	{"4", "Dezzer", 2015},
}

// GetArtist adalah handler untuk endpoint /get-artist.
func GetArtist(w http.ResponseWriter, r *http.Request) {
	// Set header Content-Type menjadi application/json.
	w.Header().Set("Content-Type", "application/json")

	// Handle request dengan method GET.
	if r.Method == "GET" {
		// Encode data artist ke format JSON.
		result, err := json.Marshal(artist)
		if err != nil {
			http.Error(w, "Error encoding data", http.StatusInternalServerError)
			return
		}
		// Kirim response JSON.
		w.Write(result)
	} else {
		// Kirim error jika method bukan GET.
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	// http.Error(w, "", http.StatusBadRequest) // Tidak perlu karena sudah dihandle di atas
}

func main() {
	// Daftarkan handler GetArtist untuk path /get-artist.
	http.HandleFunc("/get-artist", GetArtist)
	fmt.Println("Server is running on http://localhost:8080")
	// Start server HTTP pada port 8080.
	http.ListenAndServe(":8080", nil)
}
