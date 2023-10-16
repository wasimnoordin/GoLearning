// ZERO VALUES

package main

import "fmt"

/* variables declared without an explicit initial value are given their zero value

This is:
	- 0 for numeric types
	- false for booleans
	- "" (empty string) for strings

*/

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q \n", i, f, b, s)
}
