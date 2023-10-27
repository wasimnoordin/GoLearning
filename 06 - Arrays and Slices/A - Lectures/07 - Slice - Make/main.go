package main

import "fmt"

func main() {
	// using string literal
	si := []string{"a", "b", "c"}
	fmt.Println(si)

	//using make
	xi := make([]int, 0, 10) // 0 is initial store length, 10 is max capacity
	fmt.Println(xi)
	fmt.Println(len(xi)) // length
	fmt.Println(cap(xi)) // capacity

	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Println("-----------------")
	xi = append(xi, 10, 11, 12, 13)
	fmt.Println(xi)
	fmt.Println(len(xi)) // length increases after append!
	fmt.Println(cap(xi)) // capacity increases after append!
}
