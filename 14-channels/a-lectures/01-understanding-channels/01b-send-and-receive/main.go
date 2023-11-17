package main

import (
	"fmt"
)

func main() {
	// Create an unbuffered channel of type int
	c := make(chan int)

	// Start a new goroutine
	go func() {
		// Send the value 42 into the channel 'c'
		c <- 42
	}()

	// Receive the value from channel 'c'
	// This line will wait until a value is available in the channel
	fmt.Println(<-c)
}
