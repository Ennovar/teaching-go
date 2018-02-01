package main

import "fmt"

// Short declaration/initialization syntax, which automatically detects the
// type of a variable, can NOT be used in the global scope. It will generate
// a compiler error.
// myFloat := 3.14

// Declaring without initializing in the global scope requires type specification
var myString string

// Declaring and initializing in the global scope allows you to omit the type
var myInt = 10

func main() {
	fmt.Printf("Variable myInt: %v\n\n", myInt)

	myString = `This variable of type string that was declared in the global scope,
				but not initialized until in the scope of function main. also, back-ticks
				allow you to create multi-line strings\n`
	fmt.Println(myString)

	// The shorthand := notation allows you to bypass type declaration of
	// a variable.
	autoType := "This variable was automatically detected to be of type string.\n"
	fmt.Println(autoType)

	// The underscore may look kind of silly right now, but sometimes there
	// are functions are that return multiple values, and you may only care
	// about one of them, and don't need the others. The underscore is useful
	// at conserving memory in regards to situations like that.
	_ := "The underscore denotes that you do not want to store the value of the variable."

	otherTypes := []string{
		"float32", "float64",
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"byte",
		"string",
		"bool",
	}

	fmt.Println("Common *basic* data types used in Go include:")
	for _, v := range otherTypes {
		fmt.Printf("\t- %v\n", v)
	}
}