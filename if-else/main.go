package main

import "fmt"

/*
If statement di gunakan untuk menentukan blok kode yang akan di ekseskusi, jika salah gunkana if else
case sensitif harus dengan huruf kecil semua
*/
func main() {

	isActor := "Tom Hardy"

	// if else

	if isActor == "Tom Hardy" {
		fmt.Println("Good")
	} else {
		fmt.Println("Wrong Person")
	}

	// if else if else
	singer := "Morgan Wallen"
	whatSongs := "Cover me up"

	if singer == "Morgan Wallen" || whatSongs == "Whiskey Glasses" {
		fmt.Println("country man")

	} else if singer == "Morgan Wallen" && whatSongs == "Cover me up" {
		fmt.Println("country boy")
	} else {
		fmt.Println("you not a fan boy or fan man")
	}

	// nested if
	if whatSongs == "Cow Girls" {
		if whatSongs == "Last Night" {
			fmt.Println("good song ")
		} else if whatSongs == "Cover me up" {
			fmt.Println("yes")

		}
	} else {
		fmt.Println("you dont like country music")
	}

}
