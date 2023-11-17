/*

● Using goroutines, create an incrementer program
○ have a variable to hold the incrementer value
○ launch a bunch of goroutines
■ each goroutine should
● read the incrementer value
○ store it in a new variable
● yield the processor with runtime.Gosched()
● increment the new variable
● write the value in the new variable back to the incrementer
variable
● use waitgroups to wait for all of your goroutines to finish
● the above will create a race condition.
● Prove that it is a race condition by using the -race flag
● if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej

*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	// Declare and initialize the variable to hold the incrementer value
	incrementer := 0

	// Define the number of goroutines to launch
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer

			// Yield the processor to other goroutines
			runtime.Gosched()
			v++

			// Write the updated value back to the incrementer variable
			incrementer = v

			// Print the intermediate value of the incrementer
			fmt.Println("Middle Incrementer:", incrementer)

			// Notify the wait group that this goroutine has finished
			wg.Done()
		}()
	}

	// Wait for all the goroutines to finish execution
	wg.Wait()

	// Print the final value of the incrementer
	fmt.Println("Final Incrementer:", incrementer)
}

// ^^ THIS CODE WILL HAVE A RACE CONDITION SINCE WE DON'T USE MUTEX OR ATOMICITY
