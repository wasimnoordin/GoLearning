// VARIABLES

package main

import "fmt"

// A var statement declares a list of variables, i.e. function argument lists, with the type coming last
// var can be at package or function level

var c, python, java bool

const d int = 42 // Can also define a constant (parallel type)

func main() {
	var i int
	fmt.Println(i, c, python, java, d)
}
