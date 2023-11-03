/*

Using the code from the previous example, delete a record from your map. Now print the map
out using the “range” loop

*/

package main

import "fmt"

func main() {
	m := make(map[string][]string) // 1st string is key, 2nd is value
	m[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}

	// add record

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	// delete record

	delete(m, `no_dr`)

	for k, v := range m {
		fmt.Println(k)         // don't need to display value twice
		for i, v2 := range v { // range over value, as value is the slice
			fmt.Println(i, v2)
		}
	}
}
