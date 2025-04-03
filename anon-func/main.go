package main

import "fmt"

/*
Anonymous function adalah fungsi tanpa nama yang bisa langsung dieksekusi atau disimpan dalam variabel.
Anonymous function sangat berguna dalam callback, goroutine, atau kode yang hanya digunakan sekali.
*/

func main() {

	// Anonymous Function Sederhana
	func() {
		fmt.Println("Ini contoh dasar")

	}()

	// Anonymous Function dengan Parameter
	func(name string) {
		fmt.Println("Hai", name)

	}("Jhon")

	// Anonymous Function yang Disimpan dalam Variabel
	great := func(country string) string {
		return "Hai," + country
	}
	fmt.Print(great("Jhon"))

}
