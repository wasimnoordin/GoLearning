package main

import (
	"fmt"
)

func main() {
	xs := []string{"Almond Biscotti Cafe", "Banana Pudding", "Balsamic Strawberry (GF)"}
	fmt.Println(xs)
	fmt.Printf("%T \n", xs)

	// blank indetifier to not use a returned value for index
	for _, val := range xs {
		fmt.Printf("value: %v \n", val)
	}

	// accessing by index position
	fmt.Println("")

	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
	// doesn't work as slice has length 3 - 1 = 2
	// fmt.Println(xs[3])

	fmt.Println("")
	fmt.Printf("Length of the slice is: %v \n", len(xs))

	//accessing by index using a for loop

	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i])
	}
}
