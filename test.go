package main

import "fmt"

func main() {
	var x string = "x"
	var pointer *string

	pointer = &x

	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(pointer)
	fmt.Println(&pointer)
	fmt.Println(*pointer)
}
