package main

import "fmt"

type person struct {
	firstName         string
	lastName          string
	favouriteIceCream []string
}

func main() {
	p1 := person{
		firstName:         "Finn",
		lastName:          "The Human",
		favouriteIceCream: []string{"Vanilla", "Bubblegum"}, // Initializing a slice of strings.
	}

	p2 := person{
		firstName:         "Jake",
		lastName:          "The Dog",
		favouriteIceCream: []string{"Butterscotch", "Lilac and Gooseberries"},
	}

	// Print the entire struct.
	fmt.Println(p1)
	fmt.Println(p2)

	// Use a for loop with range to iterate over the slice of 'favouriteIceCream'.
	// The range keyword provides both the index and value for each entry in the slice.
	// Here, we're only interested in the value, so we use the blank identifier "_" to ignore the index.

	for _, v1 := range p1.favouriteIceCream {
		fmt.Println(p1.firstName, "favourite icecream:", v1)
	}

	for _, v2 := range p2.favouriteIceCream {
		fmt.Println(p2.firstName, "favourite icecream:", v2)
	}
}
