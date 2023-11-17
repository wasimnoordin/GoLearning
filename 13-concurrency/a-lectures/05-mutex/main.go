// A “mutex” is a mutual exclusion lock. Mutexes allow us to lock our code so that only one goroutine can access that locked chunk of code at a time
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	// Counter to track the number of operations.
	counter := 0

	// Number of goroutines to be executed.
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	// Mutex to synchronize access to the counter.
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		// Launch a goroutine.
		go func() {
			// Lock the mutex to prevent concurrent access to counter.
			mu.Lock()
			v := counter
			// Yield the processor to allow other goroutines to run.
			runtime.Gosched()
			v++
			counter = v
			// Unlock the mutex to allow other goroutines to access counter.
			mu.Unlock()
			wg.Done()
		}()
		// Print the number of active goroutines after starting each goroutine.
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
