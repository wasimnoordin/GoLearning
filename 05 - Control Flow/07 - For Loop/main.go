/*
	- there are two parts ot this hands on exercise
		-- create a program that has a loop that prints every number from 0 to 99
		-- modify the program from the previous hands on exercise to run 100 times
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i >= 0 && i < 100; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("The value of x is %v \n", x)
		fmt.Printf("The value of y is %v \n", y)

		// SWITCH STATEMENT:

		switch {
		case x < 4 && y < 4:
			fmt.Println("Both x and y are less than 4")
		case x > 6 && y > 6:
			fmt.Println("Both x and y are more than 6")
		case x >= 4 && x <= 6:
			fmt.Println("x is from 4 to 6 inclusive of both numbers")
		case y != 5:
			fmt.Println("y does not equal 5")
		default:
			fmt.Println("None of the previous cases were met!")
		}
	}
}
