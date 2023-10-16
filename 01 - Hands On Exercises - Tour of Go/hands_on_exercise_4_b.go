// VARIABLES WITH INITIALISERS

package main

import "fmt"

// var declaration can include initialisers, 1 per var
// Notice that the type is omitted (i.e. no int!)

var i, j = 1, 2

// var i, j int = 1, 2 -> Works too, but type is redundant

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
