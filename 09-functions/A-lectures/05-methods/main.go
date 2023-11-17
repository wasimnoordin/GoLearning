package main

import "fmt"

type person struct {
	firstName string
}

// Define a method named 'speak' for the 'person' type.
// Methods are functions with a special receiver argument.
func (p person) speak() {
	fmt.Println("I am", p.firstName)
}

func main() {
	p1 := person{
		firstName: "Rigby",
	}

	p2 := person{
		firstName: "Mordecai",
	}

	p1.speak()
	p2.speak()
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }
