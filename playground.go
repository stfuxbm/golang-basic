package main

import "fmt"

func main() {

	// ini adalah array
	var x = [4]string{
		"hai",
		"halo",
		"yes",
	}
	fmt.Println(x[1:3])

	// ini adalah slice
	var z = []int{
		1, 2,
	}
	fmt.Println(z[1])
	fmt.Println(z[0:1])
}
