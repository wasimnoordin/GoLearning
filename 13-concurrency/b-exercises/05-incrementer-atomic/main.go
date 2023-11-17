/*

Fix the race condition you created in exercise #4 by using package atomic

*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	// Using an int64 for atomic operations
	var incrementer int64

	gs := 100 // goroutines
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// Perform atomic read and increment operation
			atomic.AddInt64(&incrementer, 1)

			// Print the intermediate value of the incrementer
			fmt.Println("Middle Incrementer:", atomic.LoadInt64(&incrementer))

			// Notify the wait group that this goroutine has finished
			wg.Done()
		}()
	}
	wg.Wait()

	// Print the final value of the incrementer
	fmt.Println("Final Incrementer:", incrementer)
}
