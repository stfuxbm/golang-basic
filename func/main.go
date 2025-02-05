package main

import "fmt"

/*
function blok kode yang bisa di panggil berulang ulang
biasanya nama func menggunakan camelCase
parameter sebagai variable dalam func
*/
func main() {
	basicFunc()                                    // Output: Hallo
	getUserName("Halo", "Tom Hardy")               // Output: Halo Tom Hardy
	fmt.Println(getResult(1, 4, 6))                // Output: 1
	getResultWithVariable(1, 5, 6)                 // Output: 11
	fmt.Println(calculateDivisionAndRemainder(10)) // Output: 3 1
}

// basic func
func basicFunc() {
	fmt.Println("Hallo")
}

// func dengan parameter
func getUserName(message string, userName string) {
	fmt.Println(message, userName)
}

// func dengan return parameter
func getResult(a int, b int, c int) int {
	return a + b/c
}

// func named return values
func getResultWithVariable(a int, b int, c int) {
	res := a*b + c
	fmt.Println(res)
}

// func with multiple return values
func calculateDivisionAndRemainder(x int) (quotient, remainder int) {
	quotient = x / 3
	remainder = x % 3
	return
}
