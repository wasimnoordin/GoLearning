/*

Starting with this code, sort the []user by
● age
● last

Also sort each []string “Sayings” for each user
● print everything in a way that is pleasant

*/

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// ByAge defines a sorting interface for sorting users by Age
type ByAge []user

// Implementing sort.Interface for ByAge to sort users by Age
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByLast defines a sorting interface for sorting users by Last name
type ByLast []user

// Implementing sort.Interface for ByLast to sort users by Last name
func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	// Users slice
	users := []user{u1, u2, u3}

	// Printing initial unsorted users
	fmt.Println("Initial Unsorted Users:")
	printUsers(users)

	// Sorting by age
	sort.Sort(ByAge(users))
	fmt.Println("\nUsers Sorted by Age:")
	printUsers(users)

	// Sorting by last name
	sort.Sort(ByLast(users))
	fmt.Println("\nUsers Sorted by Last Name:")
	printUsers(users)
}

// Function to print user details
func printUsers(users []user) {
	for _, u := range users {
		fmt.Printf("Name: %s %s, Age: %d\nSayings:\n", u.First, u.Last, u.Age)
		for _, saying := range u.Sayings {
			fmt.Printf("\t- %s\n", saying)
		}
		fmt.Println()
	}
}
