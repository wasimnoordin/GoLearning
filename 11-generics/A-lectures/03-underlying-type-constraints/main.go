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

// myAlias is a custom type based on int.
type myAlias int

// myNums interface allows types with underlying types int or float64.
// The '~' symbol means it includes any type that is based on int or float64.
type myNums interface {
	~int | ~float64
}

// addT is a generic function summing two values of the same type.
// It accepts any type T that satisfies myNums.
func addT[T myNums](a, b T) T {
	return a + b
}

func main() {
	// Demonstrates addI and addF with their respective types.
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))

	// Demonstrates addT with a custom type (myAlias) and a float64.
	var n myAlias = 42
	fmt.Println(addT(n, 2))
	fmt.Println(addT(1.2, 2.2))
}
