package main

import (
	"fmt"
)

func main() {
	xi := []int{42, 43, 44}
	fmt.Println(xi)
	fmt.Println("---")

	// variadic parameter = append
	xi = append(xi, 45, 46, 47, 999, -1) // add new values to xi
	fmt.Println(xi)
	fmt.Println("---")
}
