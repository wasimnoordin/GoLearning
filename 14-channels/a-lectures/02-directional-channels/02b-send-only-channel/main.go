package main

import (
	"fmt"
)

func main() {
	// Create a send-only channel of type int
	cs := make(chan<- int)

	// Start a goroutine that attempts to send the value 42 into the channel 'cs'
	go func() {
		cs <- 42
	}()

	// Attempt to perform a receive operation on the send-only channel 'cs'
	// This line causes a compilation error as a receive operation is not allowed on a send-only channel
	// It's invalid to perform a receive operation on a channel that only allows sending

	// fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
