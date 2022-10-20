package main

import "fmt"

func main() {
	a := 15
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	fmt.Println(*b)
	fmt.Println(a)

	// Change val with pointer
	*b = *b + 1
	fmt.Println(a, " ", b)
}
