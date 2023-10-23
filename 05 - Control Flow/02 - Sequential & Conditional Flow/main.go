/*
create a program that uses both SEQUENTIAL and CONDITIONAL control flow. Your
program should do the following:

	- create a random int between 0 and 250
	- store the value of that int in a variable with the identifier of x
		-- func Intn(n int) int
			--- rand.Intn()
	- print out the name and value of the variable

	- use an if statement to do the following
		-- if the value is between 0 and 100
			--- print between 0 and 100
		-- if the value is between 101 and 200
			--- print between 101 and 200
		-- if the value is between 201 and 250
			--- print between 201 and 250

	- re: inclusive, exclusive – does rand.Intn()
		-- include zero in its output?
		-- include 250 in its output?
		-- show this in code using the numbers 0 - 3
	- hint:
		-- &&
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := rand.Intn(251) //250 inclusive
	fmt.Printf("The value of i is %v \n", i)

	if i >= 0 && i <= 100 {
		fmt.Println("The value of i is between 0 and 100")
	} else if i >= 101 && i <= 200 {
		fmt.Println("The value of i is between 101 and 200")
	} else if i >= 201 && i <= 250 {
		fmt.Println("The value of i is between 201 and 250")
	} else {
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
