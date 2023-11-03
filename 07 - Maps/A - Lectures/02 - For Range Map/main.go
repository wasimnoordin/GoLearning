package main

import "fmt"

func main() {
	am := make(map[string]int)
	am["Ligma"] = 28
	am["Deez"] = 37
	am["Updog"] = 99

	// for range over a map
	for k, v := range am { // k = key, v = value
		fmt.Println(k, v)
	}

	// only value
	for _, v := range am {
		fmt.Println(v)
	}

	// only key
	for k := range am {
		fmt.Println(k)
	}

	fmt.Println("")

	// recall for range with a slice

	xi := []int{423, 101, 999}

	for i, v := range xi { // i = index, v = value
		fmt.Println(i, v)
	}

	// only value
	for _, v := range xi {
		fmt.Println(v)
	}

	// only index
	for i := range xi {
		fmt.Println(i)
	}

	fmt.Println("")

}
