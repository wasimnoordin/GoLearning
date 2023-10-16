// FUNCTIONS

package main

import "fmt"

// a function can take zero or more arguments, and
// the TYPE comes AFTER the variable NAME

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
	sayHello()
}

func sayHello() {
	fmt.Println("Hello")
}
