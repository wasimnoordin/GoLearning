package main

import "fmt"

type person struct {
	firstName string
}

// 'parkEmployee' embeds 'person' struct
type parkEmployee struct {
	person
	myMom bool
}

func (p person) speak() {
	fmt.Println("I am", p.firstName)
}

func (pemp parkEmployee) speak() {

	fmt.Println("I am a park employee", pemp.firstName)
}

// Define an interface named 'human'.
// An interface is a type that specifies a behavior (in this case, the 'speak' method).

type human interface {
	// Any type that has a 'speak' method is implicitly implementing the 'human' interface.
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	pemp1 := parkEmployee{
		person: person{
			firstName: "Muscleman",
		},
		myMom: true,
	}

	p2 := person{
		firstName: "Starla",
	}

	// Both 'parkEmployee' and 'person' types implement the 'human' interface because they both have a 'speak' method.
	saySomething(pemp1)
	saySomething(p2)
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }
