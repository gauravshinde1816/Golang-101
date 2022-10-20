package main

import "fmt"

func main() {
	// MAIN DATA TYPES

	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Brad"

	var age int32 = 45
	const isCool = true
	var size float32 = 2.3

	name, email := "GauravShinde", "GauravShinde@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
