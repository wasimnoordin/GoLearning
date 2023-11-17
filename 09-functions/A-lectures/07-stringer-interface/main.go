package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

// Implement the 'String' method for the 'book' type.
// This method is used to provide a custom string representation for the type.
// It's a part of the 'Stringer' interface from the 'fmt' package.

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

// Implement the 'String' method for the 'count' type.
func (c count) String() string {

	// Convert the 'count' (which is an int) to a string using 'strconv.Itoa' and return a formatted string.
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "West With The Night",
	}

	var n count = 42

	fmt.Println(b)
	fmt.Println(n)
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }
