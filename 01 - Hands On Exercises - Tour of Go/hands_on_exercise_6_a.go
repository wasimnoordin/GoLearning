/* TYPE INFERENCE

When declaring a variable without specifying an explicit type (either := or var =), variable's type is inferred from value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:

	var i int
	j := i // j is an int

But when RHS contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on precision of the constant:

	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
*/

package main

import (
	"fmt"
)

func main() {
	v := 42 //change me! i.e. to "42"
	fmt.Printf("v is of type %T \n", v)
}
