package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel of type int with a capacity of 1
	c := make(chan int, 1)

	// Send the value 42 into the buffered channel 'c'
	// This send operation succeeds since the channel has a buffer size of 1
	c <- 42

	// Receive the value from channel 'c'
	// Print the value received (42)
	fmt.Println(<-c)

	// Attempt to send 43 into the channel 'c'
	// This send operation fails because the channel is already full (buffer size: 1)
	// It blocks until there is space available in the channel buffer, but no receiver exists for the value
	c <- 43

	// Attempt to send 44 into the channel 'c'
	// This line is reached only after the value 43 is successfully received by some receiver
	// However, it also fails due to the channel being full and blocks similarly to the previous send operation
	c <- 44

	// Receive a value from channel 'c'
	// This receives the value 43 (the first value sent after the initial successful send)
	fmt.Println(<-c)
}
