package main

import "fmt"

// addI adds two integers.
func addI(a, b int) int {
	return a + b
}

// addF adds two float64 numbers.
func addF(a, b float64) float64 {
	return a + b
}

// addT is a generic function for adding numbers.
// It works with both int and float64 due to the type constraint [T int | float64].
func addT[T int | float64](a, b T) T {
	return a + b
}

func main() {
	// Using addI for integers.
	fmt.Println(addI(1, 2))

	// Using addF for float64 numbers.
	fmt.Println(addF(1.2, 2.2))

	// Using addT for both int and float64.
	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.2, 2.2))
}
