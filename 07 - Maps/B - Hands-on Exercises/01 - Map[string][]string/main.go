/*

Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of
the values, along with their index position in the slice.

key | value
`bond_james` | `shaken, not stirred`, `martinis`, `fast cars`
`moneypenny_jenny` | `intelligence`, `literature`, `computer science`
`no_dr` | `cats`, `ice cream`, `sunsets`

*/

package main

import "fmt"

func main() {
	m := make(map[string][]string) // 1st string is key, 2nd is value
	m[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}

	// fmt.Println(m)
	// fmt.Println("")
	// fmt.Printf("%#v", m)

	for k, v := range m {
		fmt.Println(k)         // don't need to display value twice
		for i, v2 := range v { // range over value, as value is the slice
			fmt.Println(i, v2)
		}
	}
}
