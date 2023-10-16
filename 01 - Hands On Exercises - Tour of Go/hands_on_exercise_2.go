// PACKAGES, IMPORTS, & EXPORTED NAMES

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// Intn includes the lower bound but excludes the upper bound
	fmt.Println("My Favourite number is", rand.Intn(1000))

	fmt.Println("Now you have %g problems .\n", math.Sqrt(7))

	// In Go, a package is imported if it begins with a captial letter, i.e. Pi, NOT pi
	fmt.Println(math.Pi)
}
