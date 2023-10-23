/*
Modify the previous program to have your program use the init func to print
		-- "This is where initialization for my program occurs"
	- read about the init function in effective go
		-- What does niladic mean?
			--- Niladic refers to a function that takes no arguments (nil=zero, adic=arguments accepted)
*/

package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is from the init function")
}

func main() {
	i := rand.Intn(251) //250 inclusive
	fmt.Printf("The value of i is %v \n", i)

	/*

		if i >= 0 && i <= 100 {
			fmt.Println("The value of i is between 0 and 100")
		} else if i >= 101 && i <= 200 {
			fmt.Println("The value of i is between 101 and 200")
		} else if i >= 201 && i <= 250 {
			fmt.Println("The value of i is between 201 and 250")
		} else {
			fmt.Println("The value of i is more than 250")
		}

	*/

	// SWITCH STATEMENT:

	switch {
	case i >= 0 && i <= 100:
		fmt.Println("The value of i is between 0 and 100")
	case i >= 101 && i <= 200:
		fmt.Println("The value of i is between 101 and 200")
	case i >= 201 && i <= 250:
		fmt.Println("The value of i is between 201 and 250")
	default:
		fmt.Println("The value of i is more than 250")
	}

	// For the inclusive/exclusive part:

	fmt.Println(rand.Intn(3)) // exclusive of 3, then test a bunch of times to ensure 3 doesn't appear
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
}
