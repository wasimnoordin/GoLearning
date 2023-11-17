/*
Fix the race condition you created in the previous exercise by using a mutex
● it makes sense to remove runtime.Gosched()
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex // Declare a Mutex to synchronize access to the incrementer variable

	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// Lock the Mutex to ensure exclusive access to the incrementer variable
			mu.Lock()
			v := incrementer

			// doesn't make sense to pause and yield if we are locking it down for the process. since we are allowing exclusive access
			// runtime.Gosched()

			v++
			incrementer = v
			// Unlock the Mutex after modifying the incrementer variable
			mu.Unlock()

			fmt.Println("Middle Incrementer:", incrementer)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final Incrementer:", incrementer)
}
