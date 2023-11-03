package main

import "fmt"

func main() {
	foo()

	// Declare a variable and assign a function expression to it.
	// This function expression represents an anonymous function.

	x := func() {
		fmt.Println("Anonymous func ran")
	}

	y := func(s string) {
		fmt.Println("This is an anonymous func showing my name", s)
	}

	x()
	y("Sim")
}

func foo() {
	fmt.Println("Foo ran")
}

// a named function with an identifier
// func (r receive) identifier(p parameter(s)) (r return(s)) { code }

// an anonymous function
// func(p parameter(s)) (r return(s)) { code }
