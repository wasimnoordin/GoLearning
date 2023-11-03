/*

Create and use an anonymous struct with these fields:
● first string
● friends map[string]int
● favDrinks []string
Print things.

*/

package main

import "fmt"

func main() {
	// create an anonymous struct
	person := struct {
		firstName string
		friends   map[string]int
		favDrinks []string
	}{
		firstName: "Benson",
		friends: map[string]int{
			"Skips": 999,
			"Mitch": 23,
		},
		favDrinks: []string{"Passionfruit Juice", "Water"},
	}

	// print out the values
	fmt.Println("Name:", person.firstName)
	fmt.Println("Friends and their ages:", person.friends)

	for name, age := range person.friends {
		fmt.Printf("%s: %d years old \n", name, age)
	}
	fmt.Println("Favourite drinks:", person.favDrinks)
}
