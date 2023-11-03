package main

import "fmt"

func main() {

	// creating a simple map using literal method, am, and accessing a value
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}
	fmt.Println("The age of Henry was", am["Henry"])
	fmt.Println(am)          // print entire map
	fmt.Printf("%#v \n", am) // print entire map with representation in Go code

	// creating a map using make method
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Stephanie"] = 37
	fmt.Println(an)          // print entire map
	fmt.Printf("%#v \n", an) // print entire map with representation in Go code
	fmt.Println(len(an))     // length of map an
}
