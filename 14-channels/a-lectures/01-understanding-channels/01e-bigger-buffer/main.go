package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel of type int with a capacity of 2
	c := make(chan int, 2)

	// Send the value 42 into the buffered channel 'c'
	// This send operation succeeds since the channel has a buffer size of 2
	c <- 42

	// Receive the value from channel 'c'
	// Print the value received (42)
	fmt.Println(<-c)

	// Send the values 43 and 44 into the buffered channel 'c'
	// Both send operations succeed as the buffer still has space available
	c <- 43
	c <- 44

	// Receive and print the values from channel 'c'
	// The channel has enough buffered space to hold both sent values
	fmt.Println(<-c) // Prints 43
	fmt.Println(<-c) // Prints 44
}
