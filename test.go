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

	var arr = [...]int{1, 2, 3, 4, 5, 6}
	c1 := arr[0:2]
	c2 := arr[1:3]
}
