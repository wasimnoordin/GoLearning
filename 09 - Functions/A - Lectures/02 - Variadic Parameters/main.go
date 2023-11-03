package main

import "fmt"

func main() {
	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum is", x)
}

// Define a function named 'sum' that takes a variadic parameter 'ii' of type int.
// A variadic parameter allows you to pass a variable number of arguments to a function.

func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T \n", ii)

	// Initialize 'n' to store the sum.
	n := 0
	// Use a for loop with range to iterate over the slice of int
	for _, v := range ii {
		n += v
	}
	return n
}
