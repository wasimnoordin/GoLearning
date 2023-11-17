package main

import "fmt"

// Define a new type named 'person' which is a struct.
// A struct is a composite data type that groups together variables under a single name.

type person struct {
	firstname string
	lastname  string
	age       int
}

func main() {
	// Create a new 'person' named 'p1' and initialize its fields.
	p1 := person{
		firstname: "Mordecai",
		lastname:  "Bluejay",
		age:       23,
	}

	// Create another 'person' named 'p2' and initialize its fields.
	p2 := person{
		firstname: "Rigby",
		lastname:  "Racoon",
		age:       23,
	}

	// Print the entirestruct.
	fmt.Println(p1)
	fmt.Println(p2)

	// Print the type using %T and its value in a detailed manner using %#v.
	fmt.Printf("%T %#v \n", p1, p1)
	fmt.Printf("%T %#v \n", p2, p2)

	// Access and print individual fields
	fmt.Println(p1.firstname, p1.lastname, p1.age)
	fmt.Println(p2.firstname, p2.lastname, p2.age)
}
