package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

// ByAge implements sort.Interface for sorting by Age.
type ByAge []Person

// Len returns the length of the slice.
func (a ByAge) Len() int { return len(a) }

// Swap swaps the elements with indexes i and j.
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less compares ages of people at indexes i and j.
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByName implements sort.Interface for sorting by First name.
type ByName []Person

// Len, Swap, and Less methods for ByName type.
func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }

func main() {
	// Creating a slice of Person.
	p1 := Person{"Benson", 30}
	p2 := Person{"Mordecai", 23}
	p3 := Person{"Pops", 100}
	p4 := Person{"Skips", 999}
	people := []Person{p1, p2, p3, p4}

	// Sorting and printing people by age.
	fmt.Println("Before sorting by age:", people)
	sort.Sort(ByAge(people))
	fmt.Println("After sorting by age: ", people)

	fmt.Println()

	// Sorting and printing people by name.
	fmt.Println("Before sorting by name:", people)
	sort.Sort(ByName(people))
	fmt.Println("After sorting by name: ", people)
}
