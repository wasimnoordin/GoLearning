/*

● Create a type engine struct, and include this field
○ electric bool
● Create a type vehicle struct, and include these fields
■ engine
■ make
■ model
■ doors
■ color
● Create two VALUES of TYPE vehicle
○ use a composite literal
● Print out each of these values.
● Print out a single field from each of these values.

*/

package main

import "fmt"

// define the engine type
type engine struct {
	electric bool
}

// define the vehicle type
type vehicle struct {
	engine engine
	make   string
	model  string
	doors  int
	colour string
}

func main() {
	// create two values of type vehicle using composite literals
	v1 := vehicle{
		engine: engine{electric: true},
		make:   "Porsche",
		model:  "Taycan",
		doors:  4,
		colour: "Red",
	}

	v2 := vehicle{
		engine: engine{electric: false},
		make:   "Mercedes",
		model:  "G-Wagon",
		doors:  4,
		colour: "Black",
	}

	// print entire vehicle values
	fmt.Println("Vehicle 1:", v1)
	fmt.Println("Vehicle 2:", v2)

	// print single fields from vehicle values
	fmt.Println("Vehicle 1 make:", v1.make)
	fmt.Println("Vehicle 2 colour:", v2.colour)
}
