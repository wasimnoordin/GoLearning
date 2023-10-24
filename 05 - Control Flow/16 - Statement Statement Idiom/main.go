/*
	- use the "statement statement" idiom to
		-- initialize x with and random int between 0 inclusive and 5 exclusive
		-- if x is 3
			--- print "x is 3"
	- run that code 100 times
	- what's the benefit of using the "statement statement" idiom?
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	count := 1
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 { //statement statement idiom
			fmt.Printf("iteration: %v \t total count %v \t value x: %v \n", i, count, x)
			count++
		}
	}
}
