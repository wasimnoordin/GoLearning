package main

import "fmt"

func main() {
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("")

	// delete the number 4: start-4 exclusive, 5-end
	xi = append(xi[:4], xi[5:]...) // need to use ... to unfurl the slice
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("")
}
