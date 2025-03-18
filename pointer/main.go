package main

import "fmt"

/*
Pointer adalah reference atau alamat memori.
Variabel pointer berarti variabel yang berisi alamat memori suatu nilai.

Variabel biasa bisa diambil nilai pointernya, caranya dengan menambahkan tanda ampersand (&) tepat sebelum nama variabel. Metode ini disebut dengan referencing.
Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan tanda asterisk (*) tepat sebelum nama variabel. Metode ini disebut dengan dereferencing.
*/

func main() {

	x := "Tom Hardy"
	y := 021

	var fullName *string = &x
	var phoneNumber *int = &y

	fmt.Println(&x)
	fmt.Println(fullName)

	fmt.Println(&y)
	fmt.Println(phoneNumber)

	x = "Jhon"
	fmt.Println(&x)

}
