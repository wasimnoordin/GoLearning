/*
create the following variables with the following scopes:
	- package level
		- create outside of func main
		- use the
			- var keyword
			- const keyword
	- block level
		- inside func main
		- use the short declaration operator
	- use the variables in func main
*/

package main

import "fmt"

var x = 1

const y = 2

func main() {
	z := 3
	fmt.Printf("the value of x is %v and the type of x is %T\n", x, x)
	fmt.Printf("the value of y is %v and the type of y is %T\n", y, y)
	fmt.Printf("the value of z is %v and the type of z is %T\n", z, z)

}
