package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	// declare y and assign it x's address
	y := &x
	fmt.Printf("%v\t%T\n", y, y)

	// *int gives us a type, *y will give us a value by dereferencing the address
	fmt.Println(*y)

	// what if we now dereference x?
	fmt.Println(*&x)

	// now what if we want to change the value of the y, which is a dereferenced x?
	*y = 43
	fmt.Println(x)

}
