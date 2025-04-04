// Contoh penggunaan package "encoding/json" untuk melakukan encoding dan decoding data JSON di Go.
// Package ini umumnya digunakan saat bekerja dengan data JSON, seperti pada komunikasi API atau membaca/menulis file JSON.
package main

import (
	"encoding/json"
	"fmt"
)

type Users struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {
	users := `{
		"Name": "Tom",
		"Age": 12
	}`

	jsonData := []byte(users)
	var dataUsers Users

	// Decode JSON ke variabel objek struct
	err := json.Unmarshal(jsonData, &dataUsers)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("Name:", dataUsers.Name)
	fmt.Println("Age:", dataUsers.Age)

	// Encode struct ke JSON string
	encodedJSON, err := json.Marshal(dataUsers)
	// err != nill = ADA ERROR
	if err != nil {
		fmt.Println("Error encoding JSON:", err.Error()) //  Jika ada error, tampilkan pesan error
		return                                           // // Hentikan eksekusi program jika ada error
	}
	fmt.Println("\nEncode Objek ke JSON String:")
	fmt.Println(string(encodedJSON))
}
