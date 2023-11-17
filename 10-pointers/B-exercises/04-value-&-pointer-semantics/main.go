package main

import "fmt"

type person struct {
	first string
}

func main() {
	p := person{
		first: "Mordecai",
	}
	fmt.Println(p)

	p = changeName(p, "Rigby")
	fmt.Println(p)

	changeNamePointer(&p, "Benson") // dont need to assign it to p for pointer method
	fmt.Println(p)

}

func changeName(p person, s string) person { // using value
	p.first = s
	return p
}

func changeNamePointer(p *person, s string) { // using pointers, no return
	p.first = s
}

/*
Create two functions to change a field in a struct called "first" of TYPE string:
● One function will use VALUE SEMANTICS
○ this function will return some VALUE of some TYPE
● The other function will use POINTER SEMANTICS
○ this function will return nothing
● don't use methods
*/
