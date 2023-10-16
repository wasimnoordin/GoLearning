/*

This program declares two variables:
	- one variable to store a value of type int8
		-- assign to it the largest number possible, then print it
	- one variable to store a value of type uint8
		-- assign to it the largest number possible, then print it

*/

package main

import "fmt"

func main() {
	var x uint8 = 255 // unsigned largest
	var y int8 = 127  // signed largest

	fmt.Printf("%v is of type %T \n", x, x)
	fmt.Printf("%v is of type %T \n", y, y)
}
