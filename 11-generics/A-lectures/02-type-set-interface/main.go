package main

import "fmt"

// addI sums two integers.
func addI(a, b int) int {
	return a + b
}

// addF sums two float64 numbers.
func addF(a, b float64) float64 {
	return a + b
}

// myNums is a type set interface for int and float64.
type myNums interface {
	int | float64
}

// addT is a generic function that sums two values of the same type.
// T is constrained to be either int or float64 by myNums.
func addT[T myNums](a, b T) T {
	return a + b
}

func main() {
	// Demonstrates addI with int arguments.
	fmt.Println(addI(1, 2))

	// Demonstrates addF with float64 arguments.
	fmt.Println(addF(1.2, 2.2))

	// Demonstrates addT with both int and float64 arguments.
	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.2, 2.2))
}
