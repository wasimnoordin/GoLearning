// see image in lecture notes for a reference of what we have created below

// Race conditions are not good code. A race condition will give unpredictable results. We will see how to fix this race condition in the next exercise

// to check: go run -race main.go

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Print the number of CPUs and active goroutines.
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	// Counter to track the number of operations.
	counter := 0

	// Number of goroutines to be executed.
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	// Loop to create goroutines.
	for i := 0; i < gs; i++ {
		// Launch a goroutine.
		go func() {
			// Read the current value of counter.
			v := counter
			// Yield the processor to allow other goroutines to run.
			runtime.Gosched()
			// Increment the value.
			v++
			// Update the counter with the incremented value.
			counter = v
			// Notify WaitGroup about the completion of the goroutine.
			wg.Done()
		}()

		// Print the number of active goroutines after starting each goroutine.
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	// Wait for all goroutines to finish.
	wg.Wait()

	// Print the final number of active goroutines and the counter value.
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
