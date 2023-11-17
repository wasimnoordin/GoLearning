/*

● “defer” multiple functions in main
○ show that a deferred func runs after the func containing it exits.
○ determine the order in which the multiple defer funcs run

*/

// note: deferred function run via last in first out!

package main

import "fmt"

func main() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	fmt.Println("main function")
	// When the main function exits, the deferred functions will be executed in reverse order.
}

// Output:
// main function
// deferred 3
// deferred 2
// deferred 1
