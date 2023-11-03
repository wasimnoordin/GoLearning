package main

import "fmt"

func main() {
	// Create a new incrementor function 'f'.
	f := incrementor()
	// Call 'f' multiple times. With each call, the value of 'x' inside the closure increases.
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 4
	fmt.Println(f()) // 5
	fmt.Println(f()) // 6

	// Create another new incrementor function 'g'.
	g := incrementor()
	// Call 'g' multiple times. Note that 'g' has its own separate closure, so it starts from 1 again.
	fmt.Println(g()) // 1
	fmt.Println(g()) // 1
	fmt.Println(g()) // 2
	fmt.Println(g()) // 3
	fmt.Println(g()) // 4
	fmt.Println(g()) // 5
	fmt.Println(g()) // 6
}

// 'incrementor' returns a function that, when called, increments its own internal counter.
func incrementor() func() int {
	x := 0
	// Return a closure that captures and modifies the 'x' variable.
	return func() int {
		x++
		return x
	}
}
