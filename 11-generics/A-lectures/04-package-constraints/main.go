package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// addI adds two integers.
func addI(a, b int) int {
	return a + b
}

// addF adds two float64 numbers.
func addF(a, b float64) float64 {
	return a + b
}

// myNums is a type constraint interface that includes all types covered by
// constraints.Integer and constraints.Float from the constraints package.
type myNums interface {
	constraints.Integer | constraints.Float
}

// addT is a generic function that sums two values of the same type T,
// constrained by the myNums interface.
func addT[T myNums](a, b T) T {
	return a + b
}

func main() {
	// Demonstrating the use of both specific and generic add functions.
	fmt.Println(addI(1, 2))     // Adds two integers
	fmt.Println(addF(1.2, 2.2)) // Adds two float64 numbers
	fmt.Println(addT(1, 2))     // Uses generic function for integers
	fmt.Println(addT(1.2, 2.2)) // Uses generic function for float64 numbers
}
