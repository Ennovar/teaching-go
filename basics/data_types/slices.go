package main

import "fmt"

// You can declare a slice by defining its type in the
// global scope just as you would any other data type.
// The only difference with slices is the [] before
// the type. This slice will be initialized in main.
var myFavoriteNumbers []float32

func main() {
	myFavoriteNumbers = []float32{
		3.14, 12.65, 14.9, 92.43,
	}
	myFavoriteNumbers = append(myFavoriteNumbers, 43.5)

	// You can also declare and initialize a slice on the
	// fly using the shorthand := syntax. It is a little
	// different than what you would do declaring a basic
	// type, however.
	randomStrings := []string{
		"Where is the twisted proton?",
		"The energy of your resurrections will rise purely when you visualize that purpose is the aspect.",
		"Be shining for whoever occurs, because each has been needed with core.",
		"Particle of a seismic mystery, examine the life!",
	}

	// Slices can easily be iterated over using a for loop.
	// The range keyword can be looked at as a function that
	// returns a key and a value for each iterated element in
	// the slice. In this case I am discarding the key of each
	// element, as it will just be the index.
	for _, v := range randomStrings {
		fmt.Println(v)
	}
}