package main

import "fmt"

func main() {
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %#v \n", xi)
	fmt.Println("")

	// [inclusive:exclusive]
	fmt.Printf("xi - %#v \n", xi[0:4]) // get index 0,1,2,3 but not 4
	fmt.Println("")

	// [:exclusive]
	fmt.Printf("xi - %#v \n", xi[:7]) //beginning to 6
	fmt.Println("")

	// [inclusive:]
	fmt.Printf("xi - %#v \n", xi[4:]) // 4 to the end
	fmt.Println("")

	// [:]b
	fmt.Printf("xi - %#v \n", xi[:]) // everything
	fmt.Println("")

}
