package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct untuk data bahasa pemrograman
type ProgrammingLanguage struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Founder  string `json:"founder"`
	Year     int    `json:"year"`
	Paradigm string `json:"paradigm"`
}

// Struct response JSON standar
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

// Slice sebagai database sementara (di RAM)
var languages = []ProgrammingLanguage{
	{"1", "Python", "Guido van Rossum", 1991, "Object-Oriented, Procedural"},
	{"2", "C++", "Bjarne Stroustrup", 1985, "Object-Oriented, Procedural"},
	{"3", "Golang", "Griesemer, Pike, Thompson", 2009, "Concurrent, Imperative"},
	{"4", "JavaScript", "Brendan Eich", 1995, "Event-Driven, Functional"},
}

// Konstanta pesan untuk response
const (
	MsgMethodNotAllowed = "Method Not Allowed"
	MsgIDRequired       = "ID is required"
	MsgIDAlreadyExists  = "ID already exists"
	MsgIDNotFound       = "ID not found"
	MsgInvalidJSON      = "Invalid JSON"
	MsgLanguageAdded    = "Language added successfully"
	MsgLanguageFound    = "Language found"
	MsgLanguageList     = "List of programming languages"
	MsgLanguageDeleted  = "Language deleted successfully"
	MsgLanguageUpdated  = "Language updated successfully"
)

// writeJSON adalah helper function untuk mengirim response JSON dengan status code.
// Biar nggak nulis berulang-ulang di tiap handler.
func writeJSON(w http.ResponseWriter, statusCode int, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

// Handler GET semua bahasa pemrograman
func GetLanguages(w http.ResponseWriter, r *http.Request) {

	// Kalau bukan GET, kirim response error "Method Not Allowed" (405)
	if r.Method != http.MethodGet {

		writeJSON(w, http.StatusMethodNotAllowed, Response{
			false, nil, MsgMethodNotAllowed, http.StatusMethodNotAllowed,
		})
		return // untuk menghentikan logic yang nggak perlu,
	}

	// Ini ambil dari slice `languages` yang disimpan di RAM
	writeJSON(w, http.StatusOK, Response{
		true, languages, MsgLanguageList, http.StatusOK,
	})
}

func GetLanguagesById(w http.ResponseWriter, r *http.Request) {
	// Cek method HTTP, hanya izinkan GET
	if r.Method != http.MethodGet {
		// Jika method selain GET, kirim response "Method Not Allowed" (405)
		writeJSON(w, http.StatusMethodNotAllowed, Response{
			false, nil, MsgMethodNotAllowed, http.StatusMethodNotAllowed,
		})
		return // Hentikan eksekusi karena method tidak valid
	}

	// Ambil parameter "id" dari query string (contoh: /get-languages-byid?id=2)
	id := r.URL.Query().Get("id")
	// Validasi: jika ID kosong, kirim response error 400 (Bad Request)
	if id == "" {
		writeJSON(w, http.StatusBadRequest, Response{
			false, nil, MsgIDRequired, http.StatusBadRequest,
		})
		return
	}

	// Loop semua data bahasa pemrograman yang ada di slice `languages`
	for _, lang := range languages {
		// Cek apakah ID dari data sama dengan ID dari query
		if lang.ID == id {
			// Jika cocok, kirim response sukses dengan data bahasa tersebut
			writeJSON(w, http.StatusOK, Response{
				true, lang, MsgLanguageFound, http.StatusOK,
			})
			return // Hentikan proses setelah data ditemukan
		}
	}

	// Jika tidak ditemukan setelah loop, kirim response "ID Not Found" (404)
	writeJSON(w, http.StatusNotFound, Response{
		false, nil, MsgIDNotFound, http.StatusNotFound,
	})
}

func AddLanguages(w http.ResponseWriter, r *http.Request) {
	// Validasi method, hanya izinkan POST
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, Response{
			false, nil, MsgMethodNotAllowed, http.StatusMethodNotAllowed,
		})
		return // Stop eksekusi kalau bukan POST
	}

	// Buat variabel untuk nampung data JSON yang dikirim oleh user
	var newLang ProgrammingLanguage

	// Decode JSON dari body request ke struct newLang
	err := json.NewDecoder(r.Body).Decode(&newLang)
	if err != nil {
		// Kalau JSON tidak valid (misalnya salah format), kirim error 400
		writeJSON(w, http.StatusBadRequest, Response{
			false, nil, MsgInvalidJSON, http.StatusBadRequest,
		})
		return
	}

	// Validasi: Cek apakah ID sudah pernah digunakan (tidak boleh dobel)
	for _, lang := range languages {
		if lang.ID == newLang.ID {
			// Kalau ID sudah ada, tolak dengan error 400
			writeJSON(w, http.StatusBadRequest, Response{
				false, nil, MsgIDAlreadyExists, http.StatusBadRequest,
			})
			return
		}
	}

	// Jika ID aman, tambahkan data baru ke slice `languages`
	languages = append(languages, newLang)

	// Kirim response berhasil (201 Created)
	writeJSON(w, http.StatusCreated, Response{
		true, newLang, MsgLanguageAdded, http.StatusCreated,
	})
}

// Handler DELETE berdasarkan ID (pakai ?id=xxx)
func DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	// Cek dulu apakah method-nya DELETE, kalau bukan, tolak!
	if r.Method != http.MethodDelete {
		writeJSON(w, http.StatusMethodNotAllowed, Response{
			false, nil, MsgMethodNotAllowed, http.StatusMethodNotAllowed,
		})
		return
	}

	// Ambil query parameter 'id' dari URL (?id=123)
	id := r.URL.Query().Get("id")

	// Validasi: ID harus diisi, kalau kosong, kirim error
	if id == "" {
		writeJSON(w, http.StatusBadRequest, Response{
			false, nil, MsgIDRequired, http.StatusBadRequest,
		})
		return
	}

	// Loop slice `languages` untuk cari ID yang cocok
	for i, lang := range languages {
		if lang.ID == id {
			// Jika ditemukan, hapus elemen ke-i dengan slicing
			// ambil semua sebelum `i` dan setelah `i`, lalu gabung
			languages = append(languages[:i], languages[i+1:]...)

			// Kirim response sukses dengan data yang dihapus
			writeJSON(w, http.StatusOK, Response{
				true, lang, MsgLanguageDeleted, http.StatusOK,
			})
			return
		}
	}

	// Kalau ID nggak ditemukan di data, kirim not found
	writeJSON(w, http.StatusNotFound, Response{
		false, nil, MsgIDNotFound, http.StatusNotFound,
	})
}

// Handler PUT untuk update data berdasarkan ID
func UpdateLanguage(w http.ResponseWriter, r *http.Request) {
	// Cek apakah method yang dikirim adalah PUT
	if r.Method != http.MethodPut {
		// Kalau bukan, kirim response "Method Not Allowed"
		writeJSON(w, http.StatusMethodNotAllowed, Response{
			false, nil, MsgMethodNotAllowed, http.StatusMethodNotAllowed,
		})
		return
	}

	// Ambil ID dari query parameter (?id=xxx)
	id := r.URL.Query().Get("id")

	// Validasi: kalau ID-nya kosong, langsung balikin error
	if id == "" {
		writeJSON(w, http.StatusBadRequest, Response{
			false, nil, MsgIDRequired, http.StatusBadRequest,
		})
		return
	}

	// Buat variabel untuk menampung data hasil decode dari body request
	var updatedLang ProgrammingLanguage

	// Decode JSON dari body request ke dalam struct `updatedLang`
	err := json.NewDecoder(r.Body).Decode(&updatedLang)
	if err != nil {
		// Kalau gagal decode (format JSON salah), kirim error
		writeJSON(w, http.StatusBadRequest, Response{
			false, nil, MsgInvalidJSON, http.StatusBadRequest,
		})
		return
	}

	// Loop semua data yang ada di slice `languages`
	for i, lang := range languages {
		// Cek apakah ID-nya cocok dengan yang dikirim
		if lang.ID == id {
			// Kalau cocok, update data di index tersebut
			languages[i] = updatedLang

			// Kirim response sukses
			writeJSON(w, http.StatusOK, Response{
				true, updatedLang, MsgLanguageUpdated, http.StatusOK,
			})
			return
		}
	}

	// Kalau data dengan ID tersebut nggak ditemukan, kirim not found
	writeJSON(w, http.StatusNotFound, Response{
		false, nil, MsgIDNotFound, http.StatusNotFound,
	})
}

// Fungsi utama untuk jalankan server dan daftarkan endpoint
func main() {
	// Daftarkan handler untuk endpoint GET semua data
	http.HandleFunc("/get-languages", GetLanguages)

	// Daftarkan handler untuk endpoint GET berdasarkan ID (?id=xxx)
	http.HandleFunc("/get-languages-byid", GetLanguagesById)

	// Daftarkan handler untuk endpoint POST tambah data
	http.HandleFunc("/add-languages", AddLanguages)

	// Daftarkan handler untuk endpoint DELETE data berdasarkan ID (?id=xxx)
	http.HandleFunc("/delete-language", DeleteLanguage)

	// Daftarkan handler untuk endpoint PUT update data berdasarkan ID (?id=xxx)
	http.HandleFunc("/update-language", UpdateLanguage)

	// Info ke terminal kalau server sudah aktif di port 8080
	fmt.Println("✅ Server is running on http://localhost:8080")

	// Jalankan server HTTP di port 8080
	http.ListenAndServe(":8080", nil)
}
