/* BASIC TYPES

Go's basic types are:

bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32
     // represents a Unicode code point
float32 float64
complex64 complex128
*/

package main

import (
	"fmt"
	"math/cmplx"
)

// variable declarations can also be factored into blocks like import statements
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1 // just use int usually, unless you have a reason to use a sized or unsigned (positive) int type
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe) // %T is type, %v is value
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
