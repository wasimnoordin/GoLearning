/*

Starting with this code, unmarshal the JSON into a Go data structure. Hint:
https://mholt.github.io/json-to-go/


*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// To unmarshal the JSON string s into a Go data structure, first, you need to define a Go struct that matches the JSON structure.

type Person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var people []Person

	// unmarshal JSON string into slice of Person structs
	// &people allows json.Unmarshal to directly operate on the original slice

	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// print the unmarshalled slice of Person struct, don't need index
	for i, Person := range people {
		fmt.Println("Person #:", i)
		fmt.Printf("%s %s, Age: %d, Sayings: %v\n", Person.First, Person.Last, Person.Age, Person.Sayings)
		fmt.Println()
	}

}
