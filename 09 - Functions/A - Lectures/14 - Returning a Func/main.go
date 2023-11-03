package main

import "fmt"

func main() {
	// Call 'foo' to get an integer and print it.
	x := foo()
	fmt.Println(x)

	// Call 'bar' to get a function, then call that function and print its result.
	y := bar()
	fmt.Println(y())

	/*
		NB:
		- x is an integer.
		- y is a function that returns an integer.
		- y() is the result of calling that function, which is an int
	*/

	// Print the types of 'foo', 'bar', and the returned function from 'bar'.
	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", y)
}

// 'foo' returns an integer.
func foo() int {
	return 42
}

// 'bar' returns a function that returns an integer.
func bar() func() int {
	return func() int {
		return 43
	}
}
