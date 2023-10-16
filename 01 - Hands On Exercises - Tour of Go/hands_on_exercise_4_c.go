// SHORT VARIABLE DECLARATIONS

package main

import "fmt"

func main() {
	var i, j int = 1, 2

	// short assignment statement, :=, can be used instead of var declaration with implicit type (compiler will figure out type)
	// := construct is only available INSIDE a function
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
