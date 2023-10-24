/*
	- use modulus and the continue statement in a loop to print all ODD numbers
*/

package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		// if even, skip rest of loop using continue
		if i%2 == 0 {
			continue
		}
		// print odd
		fmt.Println(i)
	}
}
