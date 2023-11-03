/*

Using the code from the previous example, add a record to your map. Now print the map out
using the “range” loop
key | value
`fleming_ian` | `steaks`, `cigars`, `espionage`

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

	// add record

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range m {
		fmt.Println(k)         // don't need to display value twice
		for i, v2 := range v { // range over value, as value is the slice
			fmt.Println(i, v2)
		}
	}
}
