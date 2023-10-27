package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("")

	a[0] = 7 // lets change index pos 0 to 7 for a

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("")

	// this change occurs because both a and b are pointing to an underlying array at the same position

}
