package main

import "fmt"

// now change b so it references a different underlying array

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := make([]int, 6)
	copy(b, a) // underlying array b != underlying array a

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("")

	a[0] = 7 // lets change index pos 0 to 7 for a

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("")

	// now b will stay the same
}
