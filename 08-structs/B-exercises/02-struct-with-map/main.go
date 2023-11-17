/*

Take the code from the previous exercise, then store the VALUES of type person in a map with
the KEY of last name. Access each value in the map. Print out the values, ranging over the
slice

*/

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
		favouriteIceCream: []string{"Vanilla", "Bubblegum"},
	}

	p2 := person{
		firstName:         "Jake",
		lastName:          "The Dog",
		favouriteIceCream: []string{"Butterscotch", "Lilac and Gooseberries"},
	}

	// store the person values in a map with the key being the last name
	peopleMap := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	// range over map to access and print each person value
	for _, people := range peopleMap {
		fmt.Println(people.firstName, people.lastName)
		for _, flavour := range people.favouriteIceCream {
			fmt.Println(people.firstName, "favourite ice-cream:", flavour)
		}
		fmt.Println()
	}

}
