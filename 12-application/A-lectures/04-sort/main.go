package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 28, 26, 56, 12}
	xs := []string{"Benson", "Mordecai", "Rigby", "Skips", "Muscleman", "Hi-Five Ghost", "Pops"}

	// Display unsorted integer slice, sort it, and then display the sorted slice.
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	// Display unsorted string slice, sort it, and then display the sorted slice.
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
