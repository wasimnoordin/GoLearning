/*

● create a func with the identifier foo that
○ takes in a variadic parameter of type int
○ pass in a value of type []int into your func (unfurl the []int)
○ returns the sum of all values of type int passed in
● create a func with the identifier bar that
○ takes in a parameter of type []int
○ returns the sum of all values of type int passed in

*/

package main

import "fmt"

func foo(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func bar(i []int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func main() {
	// slice of int
	xi := []int{1, 2, 3, 4, 5}

	fooResult := foo(xi...)
	barResult := bar(xi)

	fmt.Println("foo result:", fooResult)
	fmt.Println("bar result:", barResult)
}
