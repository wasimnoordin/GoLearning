/*

● in addition to the main goroutine, launch two additional goroutines
○ each additional goroutine should print something out
● use waitgroups to make sure each goroutine finishes before your program exists

*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Initial CPUs:", runtime.NumCPU())
	fmt.Println("Initial Goroutines:", runtime.NumGoroutine())

	// Create a WaitGroup to wait for two goroutines to finish.
	var wg sync.WaitGroup
	wg.Add(2)

	// Launch first goroutine.
	go func() {
		fmt.Println("Hello, from Player 1")
		// Notify WaitGroup that this goroutine is done.
		wg.Done()
	}()

	// Launch second goroutine.
	go func() {
		fmt.Println("Hello, from Player 2")
		// Notify WaitGroup that this goroutine is done.
		wg.Done()
	}()

	fmt.Println("Intermediate CPUs:", runtime.NumCPU())
	fmt.Println("Intermediate Goroutines:", runtime.NumGoroutine())

	// Wait for both goroutines to finish.
	wg.Wait()

	fmt.Println("Exiting program...")
	fmt.Println("Final CPUs:", runtime.NumCPU())
	fmt.Println("Final Goroutines:", runtime.NumGoroutine())
}
