/*
Using the code from the previous example, use SLICING to create the following new slices
which are then printed:
● [42 43 44 45 46]
● [47 48 49 50 51]
● [44 45 46 47 48]
● [43 44 45 46 47]
https://go.dev/play/p/LoAQoGHT-W6
*/

package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// [inclusive:exclusive]
	fmt.Printf("xi - %v\n", xi[:5])
	fmt.Printf("xi - %v\n", xi[5:])
	fmt.Printf("xi - %v\n", xi[2:7])
	fmt.Printf("xi - %v\n", xi[1:6])
}
