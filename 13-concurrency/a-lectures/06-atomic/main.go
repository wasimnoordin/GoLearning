// We can use package atomic to also prevent a race condition in our incrementer code

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// Print the number of CPUs and active goroutines.
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	// counter is an int64 variable, which will be accessed atomically.
	var counter int64

	// Number of goroutines to be executed.
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		// Launch a goroutine.
		go func() {
			// Atomic operation to increment counter by 1.
			atomic.AddInt64(&counter, 1)
			// Yield the processor to allow other goroutines to run.
			runtime.Gosched()
			// Print the current value of the counter.
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			// Notify WaitGroup about the completion of the goroutine.
			wg.Done()
		}()
		// Print the number of active goroutines after starting each goroutine.
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
