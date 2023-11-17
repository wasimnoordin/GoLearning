/*

This exercise will reinforce our understanding of method sets:
● create a type person struct
● attach a method speak to type person using a pointer receiver
○ *person
● create a type human interface
○ to implicitly implement the interface, a human must have the speak method
● create func “saySomething”
○ have it take in a human as a parameter
○ have it call the speak method
● show the following in your code
○ you CAN pass a value of type *person into saySomething
○ you CANNOT pass a value of type person into saySomething
● here is a hint if you need some help
○ https://play.golang.org/p/FAwcQbNtMG

Receivers Values
-----------------------------------------------
(t T) T and *T
(t *T) *T

*/

package main

import "fmt"

// Define a person struct
type person struct {
	first string
}

// Define an interface named human
type human interface {
	speak()
}

// Method with a pointer receiver for person
func (p *person) speak() {
	fmt.Println("Wassup dawg")
}

// Function saySomething takes a human interface and calls speak method
func saySomething(h human) {
	h.speak()
}

func main() {
	// Create an instance of person
	person1 := person{
		first: "Mordecai",
	}

	// Passing a pointer of type *person to saySomething works
	saySomething(&person1)

	// This won't work since saySomething expects a pointer to person (*person)
	// saySomething(person1)

	// Calling speak method on person1 directly
	person1.speak()
}
