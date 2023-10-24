/*
	- use the code from the previous exercise
	- add this code to print a single value stored in the map
		age := m["James"]
		fmt.Println(age)
	- now use similar code to use the lookup of "Q" and print that value
	- now use the "comma ok" idiom to test whether "Q" is actually stored in the map, then
	  print out a statement if it is not stored in the map
		-- hint: check effective go for the "comma ok" idiom*/

package main

import "fmt"

func main() {

	fmt.Println("------------")
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for key, value := range m {
		fmt.Printf("key = %s, value = %v\n", key, value)
	}

	fmt.Println("------------")
	age1 := m["James"]
	fmt.Println("The age of Bond", age1)

	// check bond to see if it exists in the map using comma ok
	if value, ok := m["James"]; ok {
		fmt.Println("There is Bond, and Bond's age is", value)
	}

	age2 := m["Q"]
	fmt.Println("The age of Q", age2)

	// check Q to see if it does not exist in the map using comma ok
	if value, ok := m["Q"]; !ok {
		fmt.Println("There is no Q, and here is the zero value of an int", value)
	}
}
