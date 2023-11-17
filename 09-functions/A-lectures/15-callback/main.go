package main

import "fmt"

func main() {
	// Print the types of the 'add', 'subtract', and 'doMath' functions.
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	// Use the 'doMath' function with 'add' as the callback to sum
	// A callback is a function that's passed as an argument to another function and is executed after the completion of that function
	x := doMath(100, 33, add)
	fmt.Println(x)

	// Use the 'doMath' function with 'subtract' as the callback to subtract.
	y := doMath(100, 33, subtract)
	fmt.Println(y)
}

// 'doMath' takes two integers and a function 'f' as parameters.
// It applies the function 'f' to the two integers and returns the result.
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

// 'add' returns the sum of two integers.
func add(a int, b int) int {
	return a + b
}

// 'subtract' returns the result of subtracting the second integer from the first.
func subtract(a int, b int) int {
	return a - b
}
