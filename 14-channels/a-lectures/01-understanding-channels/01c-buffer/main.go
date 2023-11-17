package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel of type int with a capacity of 1
	c := make(chan int, 1)

	// Send the value 42 into the buffered channel 'c'
	// Since the channel has a buffer of 1, this send operation won't block immediately
	c <- 42

	// Receive the value from channel 'c'
	// This line retrieves the value from the buffered channel
	// Since there's a value in the buffer, the program proceeds without blocking
	fmt.Println(<-c)
}
