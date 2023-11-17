package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)   // gives the value
	fmt.Println(&x)  // gives the address of the value (location)
	fmt.Println(*&x) // gives value stored at a memory address
}
