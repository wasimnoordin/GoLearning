package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenney", "Gin", "Strawberry"}
	fmt.Println(jb)
	fmt.Println(jm)

	// store these now in slice of string, xxs
	// []string stores values of type string
	// [][]string stores values of type  slice of string
	xxs := [][]string{jb, jm}
	fmt.Println(xxs)
}
