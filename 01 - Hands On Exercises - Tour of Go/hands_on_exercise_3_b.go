// FUNCTIONS CONTINUED

package main

import "fmt"

// When 2 or more consecutive named function parameters share a type, you can omit it from all except the last
// I.e. add(x int, y int) becomes add(x, y int)

func add(x, y int) int {
	return x + y
}

func main() {
	a := add(100, 200)
	fmt.Println(a)
	fmt.Println(add(42, 13))
}
