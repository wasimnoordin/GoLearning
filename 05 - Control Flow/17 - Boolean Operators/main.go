/*
What do these print:
	- fmt.Println(true && true)
	- fmt.Println(true && false)
	- fmt.Println(true || true)
	- fmt.Println(true || false)
	- fmt.Println(!true)
*/

package main

import "fmt"

func main() {
	fmt.Printf("true && true \t\t%v\n", (true && true))   // prints true
	fmt.Printf("true && false \t\t%v\n", (true && false)) // prints false
	fmt.Printf("true || true \t\t%v\n", (true || true))   // prints true
	fmt.Printf("true || false \t\t%v\n", (true || false)) // prints true
	fmt.Printf("!true \t\t\t%v\n", (!true))               // prints false
}
