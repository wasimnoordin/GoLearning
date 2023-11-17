package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Create a wait group to manage goroutines.
var wg sync.WaitGroup

func main() {
	// Print system and runtime information.
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// Increment the wait group counter.
	wg.Add(1)
	// Launch a new goroutine to execute foo() concurrently.
	go foo()
	// Execute bar() in the main goroutine.
	bar()

	// Print system and runtime information after goroutine execution.
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	// Wait for all goroutines to finish executing.
	wg.Wait()
}

// foo simulates a function running concurrently in a goroutine.
func foo() {
	// Print 'foo' followed by a number in a loop.
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	// Notify the wait group that this goroutine has finished.
	wg.Done()
}

// bar simulates a function running in the main goroutine.
func bar() {
	// Print 'bar' followed by a number in a loop.
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
