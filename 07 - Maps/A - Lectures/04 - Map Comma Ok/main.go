// If we look up a non-existent key in a map, the zero value will be returned as the value associated with that key
// Comma ok can be used to check to see if you have looked up an existing key

package main

import "fmt"

func main() {
	am := make(map[string]int)
	am["Ligma"] = 28
	am["Deez"] = 37
	am["Updog"] = 99

	// // just using comma ok with if statement
	// v, ok := am["Deez"] // ok means I have found it!
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("Key does not exist")
	// }

	// comma ok is usually combined with statement statement idiom
	if val, ok := am["Deez"]; !ok {
		fmt.Println("Key does not exist")
	} else {
		fmt.Println("The value prints", val)
	}
}
