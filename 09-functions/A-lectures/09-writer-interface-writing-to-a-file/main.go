package main

import (
	"log"
	"os"
)

func main() {
	// Use 'os.Create' to create a file named 'output.txt'.
	// This function returns a file pointer 'f' and an error 'err'.
	f, err := os.Create("output.txt")

	// Check if there was an error while creating the file.
	if err != nil {
		log.Fatalf("error %s", err)
	}

	// defer to close file f after other operations in 'main'
	defer f.Close()

	// Define a slice of bytes 's' with the content "Hello gophers!".
	s := []byte("Hello gophers!")

	// Use the 'Write' method of the file pointer 'f' to write the byte slice 's' to the file.
	_, err = f.Write(s)

	// Check if there was an error while writing to the file.
	if err != nil {
		log.Fatalf("error %s", err)
	}
}
