package main

import "fmt"

/*
Map adalah tipe data berbentuk key-value pair, di mana setiap value memiliki key unik.
Key digunakan untuk mengakses, mengubah, atau menghapus value yang terkait dengannya.
*/

func main() {

	// map dengan keyword short hand
	actor := map[string]string{
		"Name": "Tom Hardy",
		"Jobs": "Actor",
	}

	movie := map[int]string{
		2024: "Venom",
		2015: "Legend",
	}

	fmt.Println(actor)
	fmt.Println(actor["Name"]) // ambil data menggunkan key
	fmt.Println(movie)
	fmt.Println(movie[2015])

	// map dengan function make
	netWorth := make(map[int]string)
	netWorth[2024] = "$55 million"

	fmt.Println(netWorth[2024])

}
