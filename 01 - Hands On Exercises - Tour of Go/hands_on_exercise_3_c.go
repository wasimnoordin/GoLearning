// MULTIPLE RESULTS

package main

import "fmt"

// A function can return any number of strings
// the "swap" function returns 2 strings

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
