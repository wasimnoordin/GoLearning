package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []float64{3, 1, 4, 2}

	fmt.Println(medianOne(n))          // call function medianOne
	fmt.Println("after medianOne:", n) // this sorts the array

	m := []float64{3, 1, 4, 2}
	fmt.Println(medianTwo(m))
	fmt.Println("after medianTwo:", m)
}

// uses the same underlying array
// everything is pass by value in go
// the value is being passed into thie func
// and then assigned to x

func medianOne(x []float64) float64 { // takes in a slice of float64
	sort.Float64s(x)
	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i/2]
	}
	return (x[i-1] + x[i]) / 2
}

func medianTwo(m []float64) float64 {
	// allocate a new underlying array
	// this means we don't need to change the underlying array of what was passed in
	n := make([]float64, len(m))
	copy(n, m)

	sort.Float64s(n)
	i := len(n) / 2
	if len(n)%2 == 1 {
		return n[i/2]
		// when you divide
		// you get the whole number quotient
		// without the remainder modulus
		// https://go.dev/play/p/2r5WrelUEh7
	}
	return (n[i-1] + n[i]) / 2
}
