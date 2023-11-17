package main

import "fmt"

func main() {
	x := 42
	fmt.Print(x)  // points to value of x
	fmt.Print(&x) // points to address of x
}

/*

- Create a value and assign it to a variable.
● Print the address of that value.

*/
