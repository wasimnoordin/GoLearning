package main

import "fmt"

func main() {

	// array literal
	a := [3]int{42, 43, 44} // 3 is number of values stored
	fmt.Println(a)

	b := [...]string{"Hello", "Gophers!"} // ... lets compiler figure out how many values are there
	fmt.Println(b)

	// array literal - declare variable first
	var c [2]int
	c[0] = 7 // position 0
	c[1] = 8 // position 1
	fmt.Println(c)

	fmt.Printf("%T\n", a) // type of array a
	fmt.Printf("%T\n", c) // type of array c

	// c = a means we store a in c, which doesn't work as a stores 3 and c stores 2

	{
		// declare a variable of type [7]int
		var ni [7]int
		// assign a value to index position zero
		ni[0] = 42
		fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)

		// array literal
		ni2 := [4]int{55, 56, 57, 58}
		fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)

		// array literal
		// have compiler count elements
		ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
		fmt.Printf("%#v \t %T\n", ns, ns)

		// use the builtin len function
		// https://pkg.go.dev/builtin#len
		fmt.Println(len(ni))
		fmt.Println(len(ni2))
		fmt.Println(len(ns))
	}
}

/*
"Arrays have their place, but they’re a bit inflexible,
so you don’t see them too often in Go code.
Slices, though, are everywhere.
They build on arrays to provide
great power and convenience."
~ Go Slices: usage and internals - Go Blog - Andrew Gerrand
*/
// https://go.dev/blog/slices-intro

// PRACTICE - next video
/*
Use a composite literal array
to store these elements in an array
and let the compiler determine the length of the array
*/
