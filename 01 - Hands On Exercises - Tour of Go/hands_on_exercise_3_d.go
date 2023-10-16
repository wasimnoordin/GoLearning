// NAMED RETURN VALUES

package main

import "fmt"

// Go's return variables may be named, if so they are defined at the top of the function
// A return statement without arguments returns the namwed return values ("naked" return)
// Naked returns should only be used in short functions, may harm readability in longer ones

// Here we are returning 2 variables x and y

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
