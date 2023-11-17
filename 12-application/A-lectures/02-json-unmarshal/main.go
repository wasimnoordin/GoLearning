// now lets unmarshal our JSON from the previous exercise

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Person struct should match the JSON structure.
type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	s := `[{"First":"Mitch","Last":"Sorenstein","Age":23},{"First":"Benson","Last":"The Boss","Age":30}]`

	var people []Person

	// Unmarshal the JSON string into the slice of Person structs.
	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Print the unmarshalled slice of Person structs.
	fmt.Println(people)
}
