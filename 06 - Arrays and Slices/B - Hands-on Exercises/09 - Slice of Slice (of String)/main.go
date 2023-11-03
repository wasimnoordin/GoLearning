/*

Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."
Range over the records, then range over the data in each record.

*/

package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Shaken, not Stirred."}
	mp := []string{"Miss", "Moneypenny", "I'm 008."}
	xxs := [][]string{jb, mp}

	for i, v := range xxs {
		fmt.Printf("Index: %v Value: %v \n", i, v)
		for a, b := range v {
			fmt.Printf("Index: %v  Value: %v \n", a, b)
		}
	}
}
