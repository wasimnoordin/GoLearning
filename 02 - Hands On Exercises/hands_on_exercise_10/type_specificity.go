/* This file accomplishes:

- (a) var for zero value
- (b) short declaration operator
- (c) multiple initialisations
- (d) var when specificity is requried
- (e) blank identifier

*/

package main

import "fmt"

var y int // declare a variable y which will store values of type int

func main() {
	fmt.Println(y) // (a)

	z := 42 // (b)
	fmt.Println(z)

	a, b := 43, "Happiness" // (c)
	fmt.Println(a, b)

	var c float32 = 42.42 // (d)
	fmt.Printf("%v is of this type %T \n", c, c)

	e, f, _ := 45, 46, 47 // (e)
	fmt.Println(e, f)

}
