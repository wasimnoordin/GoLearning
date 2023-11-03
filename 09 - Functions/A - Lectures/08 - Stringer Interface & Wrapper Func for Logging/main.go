package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

// Define a function named 'writeLog' that takes a parameter of type 'fmt.Stringer'.
// This function logs the string representation of the passed value.
func writeLog(s fmt.Stringer) {
	log.Println(s.String()) // Log the string representation of 's'.
}

func main() {
	b := book{
		title: "West With The Night",
	}

	var c count = 42

	writeLog(b)
	writeLog(c)
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }
