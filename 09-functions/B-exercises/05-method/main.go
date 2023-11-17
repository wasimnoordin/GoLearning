/*

Create a user defined struct with
○ the identifier “person”
○ the fields:
■ first
■ age
● attach a method to type person with
○ the identifier “speak”
○ the method should have the person say their name and age
● create a value of type person
● call the method from the value of type person

*/

package main

import "fmt"

type person struct {
	first string
	age   int
}

// attach method to type person

func (p person) speak() {
	fmt.Printf("Hi, I'm %s and I'm %d years old", p.first, p.age)
}

func main() {
	// create value of type person
	p := person{
		first: "Mitch",
		age:   23,
	}

	// call the method from the value of type person
	p.speak()
}
