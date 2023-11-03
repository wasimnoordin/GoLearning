package main

import "fmt"

// Define a new type named 'person' which is a struct.
type person struct {
	firstname string
	lastname  string
	age       int
}

// Define a new type named 'based' which is a struct.
// This struct embeds the 'person' struct and has additional fields.
type based struct {
	person       // Embedded 'person' struct.
	areTheyBased bool
	firstname    string // Field 'firstname' of type string. This shadows the 'firstname' from the embedded 'person' struct.
}

func main() {
	// Create a new 'based' named 'bsd1' and initialize its fields.
	bsd1 := based{
		person: person{ // Initialize the embedded 'person' struct.
			firstname: "Mordecai",
			lastname:  "Bluejay",
			age:       23,
		},
		firstname:    "Mitch", // Initialize the 'firstname' of 'based', not the embedded 'person'.
		areTheyBased: true,
	}

	p2 := person{
		firstname: "Rigby",
		lastname:  "Racoon",
		age:       23,
	}

	// Print the entire struct.
	fmt.Println(bsd1)
	fmt.Println(p2)

	// Print the type of 'bsd1' using %T and its value in a detailed manner using %#v.
	fmt.Printf("%T %#v \n", bsd1, bsd1)

	// Access and print individual fields of 'bsd1'. Note that 'bsd1.firstname' refers to the 'firstname' of 'based', not the embedded 'person'.
	fmt.Println(bsd1.firstname, bsd1.lastname, bsd1.age)

	// Access and print all fields of 'bsd1', including the embedded 'person' struct.
	fmt.Println(bsd1.firstname, bsd1.lastname, bsd1.age, bsd1.areTheyBased, bsd1.person)

	// Demonstrate the shadowing of the 'firstname' field. 'bsd1.firstname' refers to 'based' while 'bsd1.person.firstname' refers to the embedded 'person'.
	fmt.Println(bsd1.firstname, bsd1.person.firstname)
}
