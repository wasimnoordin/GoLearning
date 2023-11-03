package main

import "fmt"

func main() {
	// Calculate 4! (4 factorial) in two ways: directly and using the recursive and iterative functions.

	// Direct calculation of 4!
	fmt.Println(4 * 3 * 2 * 1)

	// Using the recursive 'factorial' function.
	fmt.Println(factorial(4))

	// Using the iterative 'factLoop' function.
	fmt.Println(factLoop(4))
}

// Recursive function to calculate the factorial of 'n'.
func factorial(n int) int {
	fmt.Println("This is n", n)
	// Base case: if n is 0, return 1.
	if n == 0 {
		return 1
	}
	// Recursive case: n times factorial of (n-1).
	return n * factorial(n-1)
}

/*
factorial(4)		--> 4 *
factorial(4-1)		--> 3 *
factorial(3-1)		--> 2 *
factorial(2-1)		--> 1 *
factorial(1-1)		--> 0 // base case
*/

// Iterative function to calculate the factorial of 'n'.
func factLoop(n int) int {
	total := 1
	// Multiply 'total' by 'n', then decrement 'n' until 'n' is 0.
	for n > 0 {
		total *= n
		n--
	}
	return total
}
