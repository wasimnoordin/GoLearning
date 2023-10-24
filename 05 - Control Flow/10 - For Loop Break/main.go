/*
	- create a for loop that uses break statement
	- increment or decrement the variable being checked as a condition in the loop
*/

package main

import "fmt"

func main() {

	i := 42
	for {
		if i < 0 {
			break
		}
		fmt.Println(i)
		i--
	}

	j := -18
	for {
		if j > 0 {
			break
		}
		fmt.Println(j)
		j++
	}
}
