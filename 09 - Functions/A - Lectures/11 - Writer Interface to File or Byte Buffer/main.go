package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	firstName string
}

func (p person) writeOut(w io.Writer) {
	// Convert the 'firstName' to a byte slice and write it to the provided writer 'w'.
	w.Write([]byte(p.firstName))
}

func main() {

	p := person{
		firstName: "Pops",
	}

	// Use the 'os.Create' function to create 'output.txt'.
	// This function returns a file pointer 'f' and an error 'err'.
	f, err := os.Create("output.txt")

	if err != nil {
		// If there's an error, log it and terminate the program.
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	// Create a new bytes buffer 'b'.
	var b bytes.Buffer
	p.writeOut(f)

	// Call the 'writeOut' method on 'p' to write its 'firstName' to byte buffer 'b'. & before b is used as a pointer to b's address
	p.writeOut(&b)
	fmt.Println(b.String())
}
