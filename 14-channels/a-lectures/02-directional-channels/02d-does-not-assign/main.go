package main

import (
	"fmt"
)

func main() {
	// Create a bidirectional channel that allows both sending and receiving of type int
	c := make(chan int)

	// Create a receive-only channel of type int
	cr := make(<-chan int) // receive

	// Create a send-only channel of type int
	cs := make(chan<- int) // send

	fmt.Println("-----")
	// Print the types of the channels created
	fmt.Printf("c\t%T\n", c)   // Type of 'c' is a bidirectional channel that allows both sending and receiving
	fmt.Printf("cr\t%T\n", cr) // Type of 'cr' is a receive-only channel
	fmt.Printf("cs\t%T\n", cs) // Type of 'cs' is a send-only channel

	// Assigning a receive-only channel 'cr' to a bidirectional channel 'c' is allowed
	// c = cr
	// Assigning a send-only channel 'cs' to a bidirectional channel 'c' is allowed
	// c = cs

	// Assigning a receive-only channel 'cr' to a send-only channel 'cs' is not allowed and causes a compilation error
	// cs = cr
}
