package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	fmt.Println("Hello Gophers")

	// test some of the functions in puppy package

	var1 := puppy.Bark()
	var2 := puppy.Barks()
	var3 := puppy.BigBark()
	var4 := puppy.BigBarks()

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)

	puppy.From12() // now we cant run From13 since v1.2.0 does not include it
}
