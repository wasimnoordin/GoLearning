package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Ensure struct fields are exported by starting them with capital letters.
type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "Mitch",
		Last:  "Sorenstein",
		Age:   23,
	}

	p2 := Person{
		First: "Benson",
		Last:  "The Boss",
		Age:   30,
	}

	people := []Person{
		p1,
		p2,
	}

	// Marshal the slice of Person structs into JSON
	jsonData, err := json.Marshal(people)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// Print the JSON string
	fmt.Println(string(jsonData))
}
