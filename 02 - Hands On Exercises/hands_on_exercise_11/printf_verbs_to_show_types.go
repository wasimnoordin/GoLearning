/*
This program achieves the following

	a. for a variable storing a value of type
		- string
		- int
		- float64
	b. use print verbs to show
		- value
		- type

*/

package main

import "fmt"

func main() {
	s, i, f := "Gopher", 42, 42.42
	fmt.Printf("%v is of type %T\n", s, s)
	fmt.Printf("%v is of type %T\n", i, i)
	fmt.Printf("%v is of type %T\n", f, f)
}
