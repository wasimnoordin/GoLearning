package main

import (
	"fmt"
)

// doSomething multiplies the input integer x by 5 and returns the result.
func doSomething(x int) int {
	return x * 5
}

func main() {
	// Create an unbuffered channel of integers.
	ch := make(chan int)

	// Launch a goroutine (anonymous function) that performs an operation and sends the result to the channel.
	go func() {
		// Calls the doSomething function and sends its result (5 * 5 = 25) to the channel.
		ch <- doSomething(5)
	}()

	// Retrieve the value from the channel and print it.
	fmt.Println(<-ch)
}
