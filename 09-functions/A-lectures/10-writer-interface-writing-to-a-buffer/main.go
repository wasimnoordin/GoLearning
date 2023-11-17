package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())

	// Append the string "Gophers!" to the buffer 'b'.
	b.WriteString("Gophers!")
	fmt.Println(b.String())

	// Reset the buffer 'b', removing all its content.
	b.Reset()

	// Write the string to the buffer 'b'.
	b.WriteString("It's time to d-d-d-d-duel... ")
	fmt.Println(b.String())

	// Write the byte slice representing the string to the buffer 'b'.
	b.Write([]byte("Heart of the cards"))
	fmt.Println(b.String())
}
