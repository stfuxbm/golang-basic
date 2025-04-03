/*
Reflect: library untuk memeriksa dan memanipulasi tipe data saat runtime.
Digunakan saat tipe data tidak diketahui pada waktu kompilasi,
misalnya saat bekerja dengan data generik atau metadata.
Memungkinkan introspeksi struktur dan tipe variabel.
*/
package main

import (
	"fmt"
	"reflect"
)

// Struktur contoh.
type People struct {
	id   string
	name string
}

// Periksa apakah interface adalah struct dan cetak namanya.
func CheckStruct(i any) {
	t := reflect.TypeOf(i) // Dapatkan tipe dari interface.

	if t.Kind() == reflect.Struct { // Cek apakah tipenya struct.
		fmt.Println("Struct Name:", t.Name()) // Cetak nama struct.
	} else {
		fmt.Println("Kind:", t.Kind()) // Cetak jenis tipe data lain.
	}
}

func main() {
	x := 30
	checkReflect := reflect.ValueOf(x) // Dapatkan nilai reflect dari x.
	fmt.Println(checkReflect)
	fmt.Println("Tipe dari x:", checkReflect.Type()) // Cetak tipe data x.

	z := People{id: "13", name: "Jhon"} // Buat instance struct People.

	checkStructValue := reflect.ValueOf(z) // Dapatkan nilai reflect dari z.
	fmt.Println("Struct Data:", z)
	fmt.Println("Struct Type:", checkStructValue.Type()) // Cetak tipe data z.

	CheckStruct(z) // Panggil fungsi untuk memeriksa struct z.
}
