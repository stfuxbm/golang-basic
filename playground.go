package main

import "fmt"

func main() {
	p := 4

	var checkOpt = p > 2 && p < 10 || p == 10
	fmt.Println(checkOpt)

}
