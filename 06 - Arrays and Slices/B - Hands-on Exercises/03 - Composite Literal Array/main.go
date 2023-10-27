/*

Using a COMPOSITE LITERAL:
	- create an ARRAY which holds 5 VALUES of TYPE int
	- assign VALUES to each index position.
	- Range over the array and print the values out.
		-- print out the VALUE and the TYPE

*/

package main

import "fmt"

func main() {
	// composite literal, create array
	xi := [5]int{}

	// assign values to each index
	for i := 0; i < 5; i++ {
		xi[i] = i
	}

	// print out
	for i, v := range xi {
		fmt.Printf("%v - %T - %v \n", v, v, i)
	}
}
