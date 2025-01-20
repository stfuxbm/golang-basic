// && = true jika kedua nilai sama ( and)
// || = true jika salah satu benar ( or)
// ! = kebalikan dari hasil true jika hasilnya true (not)

package main

import "fmt"

func main() {
	x := 12

	var checkAnd = x < 10 && x > 10
	var checkOr = x < 10 || x > 10
	var chekNotSame = !(x > 10)

	fmt.Println(checkAnd)    // false
	fmt.Println(checkOr)     // true
	fmt.Println(chekNotSame) //false
}
