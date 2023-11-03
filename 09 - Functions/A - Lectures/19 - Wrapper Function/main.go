package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Call the 'readFile' wrapper function to read the contents of "poem.txt".
	xb, err := readFile("poem.txt")

	// If there's an error reading the file, log it and terminate the program.
	if err != nil {
		log.Fatalf("error in main in readFile: %s", err)
	}

	// Print the raw bytes of the file contents.
	fmt.Println(xb)

	// Convert the bytes to a string and print the file contents.
	fmt.Println(string(xb))
}

// 'readFile' is a wrapper function around 'os.ReadFile' that reads the contents of a file.
// It returns the file contents as a byte slice and any error encountered.
func readFile(fileName string) ([]byte, error) {
	xb, err := os.ReadFile(fileName)

	// If there's an error reading the file, wrap the error with a custom message and return.
	if err != nil {
		return nil, fmt.Errorf("there was an error in readFile: %s", err)
	}

	// Return the file contents and nil (indicating no error).
	return xb, nil
}
