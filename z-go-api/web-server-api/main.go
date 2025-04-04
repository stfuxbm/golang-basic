// Web API sederhana menggunakan Go untuk CRUD data bahasa pemrograman.
// Data disimpan di RAM (slice), cocok untuk belajar dasar-dasar HTTP handler dan JSON API.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct untuk representasi data bahasa pemrograman
type ProgrammingLanguage struct {
	ID       string `json:"ID"`
	Name     string `json:"Name"`
	Founder  string `json:"Founder"`
	Year     int    `json:"Year"`
	Paradigm string `json:"Paradigm"`
}

// Struct standar untuk format response JSON
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

// Slice sebagai "database" sementara di RAM
var languages = []ProgrammingLanguage{
	{"1", "Python", "Guido van Rossum", 1991, "Object-Oriented, Procedural"},
	{"2", "C++", "Bjarne Stroustrup", 1985, "Object-Oriented, Procedural"},
	{"3", "Golang", "Griesemer, Pike, Thompson", 2009, "Concurrent, Imperative"},
	{"4", "JavaScript", "Brendan Eich", 1995, "Event-Driven, Functional"},
}

// Handler untuk GET semua bahasa pemrograman
// Digunakan saat ingin menampilkan semua data
func GetLanguages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// Kirim seluruh data language sebagai JSON
		res := Response{true, languages, "List of programming languages", http.StatusOK}
		json.NewEncoder(w).Encode(res)
	} else {
		// Method selain GET tidak diizinkan
		res := Response{false, nil, "Method Not Allowed", http.StatusMethodNotAllowed}
		json.NewEncoder(w).Encode(res)
	}
}

// Handler untuk GET data berdasarkan ID
// Digunakan saat ingin melihat detail 1 data berdasarkan ID
func GetLanguagesById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		res := Response{false, nil, "Method Not Allowed", http.StatusMethodNotAllowed}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Ambil ID dari query string
	id := r.URL.Query().Get("id")

	// Cari ID di slice languages
	for _, lang := range languages {
		if lang.ID == id {
			// Kirim data jika ditemukan
			res := Response{true, lang, "Language Found", http.StatusOK}
			json.NewEncoder(w).Encode(res)
			return
		}
	}

	// Jika tidak ditemukan
	res := Response{false, nil, "ID Not Found", http.StatusNotFound}
	json.NewEncoder(w).Encode(res)
}

// Handler untuk POST tambah data
// Digunakan saat ingin menambahkan bahasa pemrograman baru
func AddLanguages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode body JSON menjadi struct
	var newLang ProgrammingLanguage
	err := json.NewDecoder(r.Body).Decode(&newLang)
	if err != nil {
		http.Error(w, "JSON not valid", http.StatusBadRequest)
		return
	}

	// Tambahkan data ke slice
	languages = append(languages, newLang)

	// Kirim response berhasil
	res := Response{true, newLang, "Success add Data Language", http.StatusCreated}
	json.NewEncoder(w).Encode(res)
}

// Handler untuk DELETE data berdasarkan ID
// Digunakan saat ingin menghapus data bahasa pemrograman berdasarkan ID
func DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "DELETE" {
		res := Response{false, nil, "Method Not Allowed", http.StatusMethodNotAllowed}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Ambil ID dari query string
	id := r.URL.Query().Get("id")
	if id == "" {
		res := Response{false, nil, "ID is required", http.StatusBadRequest}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Cari dan hapus data dari slice
	for i, lang := range languages {
		if lang.ID == id {
			languages = append(languages[:i], languages[i+1:]...)
			res := Response{true, lang, "Language deleted successfully", http.StatusOK}
			json.NewEncoder(w).Encode(res)
			return
		}
	}

	// Jika ID tidak ditemukan
	res := Response{false, nil, "ID not found", http.StatusNotFound}
	json.NewEncoder(w).Encode(res)
}

// Handler untuk UPDATE data berdasarkan ID (PUT)
// Digunakan saat ingin mengedit data bahasa pemrograman
func UpdateLanguage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Validasi method harus PUT
	if r.Method != "PUT" {
		res := Response{false, nil, "Method Not Allowed", http.StatusMethodNotAllowed}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Ambil ID dari query
	id := r.URL.Query().Get("id")
	if id == "" {
		res := Response{false, nil, "ID is required", http.StatusBadRequest}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Decode body ke struct baru
	var updatedLang ProgrammingLanguage
	err := json.NewDecoder(r.Body).Decode(&updatedLang)
	if err != nil {
		res := Response{false, nil, "Invalid JSON", http.StatusBadRequest}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Cari dan update data
	for i, lang := range languages {
		if lang.ID == id {
			// Update hanya field yang dikirim
			languages[i] = updatedLang // langsung ganti semua field

			res := Response{true, updatedLang, "Language updated successfully", http.StatusOK}
			json.NewEncoder(w).Encode(res)
			return
		}
	}

	// Jika ID tidak ditemukan
	res := Response{false, nil, "ID not found", http.StatusNotFound}
	json.NewEncoder(w).Encode(res)
}

// Fungsi utama untuk menjalankan server dan daftarkan endpoint
func main() {
	http.HandleFunc("/get-languages", GetLanguages)
	http.HandleFunc("/get-languages-byid", GetLanguagesById)
	http.HandleFunc("/add-languages", AddLanguages)
	http.HandleFunc("/delete-language", DeleteLanguage)
	http.HandleFunc("/update-language", UpdateLanguage)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
