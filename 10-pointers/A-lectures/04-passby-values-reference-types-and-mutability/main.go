package main

import "fmt"

// intDelta sets the value pointed by n to 43.
func intDelta(n *int) {
	*n = 43
}

// sliceDelta modifies the first element of slice ii to 99.
func sliceDelta(ii []int) {
	ii[0] = 99
}

// mapDelta sets the value of key s in map md to 24.
func mapDelta(md map[string]int, s string) {
	md[s] = 24
}

func main() {
	// Demonstrating how intDelta affects an integer.
	a := 42
	fmt.Println(a) // Print initial value of a.
	intDelta(&a)
	fmt.Println(a) // Print new value of a after intDelta.

	// Demonstrating how sliceDelta affects a slice.
	xi := []int{1, 2, 3, 4}
	fmt.Println(xi) // Print initial state of slice xi.
	sliceDelta(xi)
	fmt.Println(xi) // Print new state of xi after sliceDelta.

	// Demonstrating how mapDelta affects a map.
	m := make(map[string]int)
	m["Mitch"] = 23
	fmt.Println(m["Mitch"]) // Print initial value associated with "Mitch".
	mapDelta(m, "Mitch")
	fmt.Println(m["Mitch"]) // Print new value associated with "Mitch" after mapDelta.
}
