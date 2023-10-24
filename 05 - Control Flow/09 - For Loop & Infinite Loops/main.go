/*
	- create a for loop using only a condition
	- increment or decrement the variable being checked in the condition
*/

package main

import "fmt"

func main() {
	i := 10

	for i >= 0 {
		fmt.Println(i)
		i--
	}

	j := 0
	for j <= 10 {
		fmt.Println(j)
		j++
	}

}
