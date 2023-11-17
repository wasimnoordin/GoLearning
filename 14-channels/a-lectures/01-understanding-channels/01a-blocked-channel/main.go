package main

import (
	"fmt"
)

func main() {
	// Create an unbuffered channel of type int
	ca := make(chan int)

	// This line causes a deadlock because sending to an unbuffered channel
	// blocks until there is a corresponding receiver for the value.
	// However, there's no active receiver yet, so this operation blocks
	// the main goroutine indefinitely.
	ca <- 42

	// This line is never reached due to the deadlock above.
	// The program is unable to proceed beyond the blocked send operation,
	// resulting in a deadlock situation.
	fmt.Println(<-ca)
}
