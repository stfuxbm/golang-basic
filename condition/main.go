package main

import "fmt"

func main() {
	fmt.Println("input your number: ")

	// var x int
	// fmt.Scanln(&x)
	// fmt.Println("your number is", x)
	// // use if else
	// if x == 2 {
	// 	fmt.Println("nilai tidak sama")
	// } else {
	// 	fmt.Println("faf")
	// }
	// temporary if else
	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
	// switch

}
