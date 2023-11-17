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

	fmt.Println("-----")
	// Conversion from more general to specific channel types is allowed
	// Converting a bidirectional channel 'c' to a receive-only channel '(<-chan int)(c)' is allowed
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	// Converting a bidirectional channel 'c' to a send-only channel '(chan<- int)(c)' is allowed
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	// Conversion from specific to general channel types isn't allowed and will cause a compilation error
	// These lines are commented out as they cause compilation errors
	// fmt.Println("-----")
	// fmt.Printf("c\t%T\n", (chan int)(cr)) // Converting a receive-only channel 'cr' to a bidirectional channel isn't allowed
	// fmt.Printf("c\t%T\n", (chan int)(cs)) // Converting a send-only channel 'cs' to a bidirectional channel isn't allowed
}
