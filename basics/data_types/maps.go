package main

import "fmt"

// Declaring maps in the global scope is easy,
// it follows the same structure as any other
// variable declared in the global scope.
var pageHits map[string]int

func main() {
	pageHits = map[string]int{
		"home": 43,
		"blog": 16,
	}
	pageHits["contact"] = 64

	// Making a map on the fly is a little strange
	myContacts := make(map[string]string)
	myContacts["Jim"] = "316-555-5555"
	myContacts["Ashley"] = "644-555-5555"

	// In this case of iteration, unlike how the slice
	// example was dealt with, the key and value, k and
	// v respectively, are necessary information.
	for k, v := range myContacts {
		fmt.Printf("%v's number is %v.\n", k, v)
	}
}
