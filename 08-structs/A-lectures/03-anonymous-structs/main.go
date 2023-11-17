package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
}

func main() {
	// Create an instance of an anonymous struct.
	// Anonymous structs are useful when you want a one-off struct that won't be reused elsewhere in your code.
	p1 := struct {
		firstname string
		lastname  string
		age       int
	}{
		firstname: "Mordecai",
		lastname:  "Bluejay",
		age:       23,
	}

	// Create a new 'person' named 'p2' and initialize its fields.
	p2 := person{
		firstname: "Rigby",
		lastname:  "Racoon",
		age:       23,
	}

	// Print the entire anonymous struct 'p1' and the struct 'p2'.
	fmt.Println(p1)
	fmt.Println(p2)

	// Print the type using %T.
	// This will show a detailed type signature since it's an anonymous struct.
	fmt.Printf("%T \n", p1)
	fmt.Printf("%T \n", p2)
}
