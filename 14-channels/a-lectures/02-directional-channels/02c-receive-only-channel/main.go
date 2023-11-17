package main

import (
	"fmt"
)

func main() {
	// Create a receive-only channel of type int
	cr := make(<-chan int)

	// Start a goroutine that attempts to send the value 42 into the receive-only channel 'cr'
	// This line causes a compilation error as sending to a receive-only channel is disallowed

	go func() {
		// cr <- 42
	}()

	// Attempt to perform a receive operation on the receive-only channel 'cr'
	// This line causes a compilation error as a receive operation on an uninitialized channel is disallowed
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
