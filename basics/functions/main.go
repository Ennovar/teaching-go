package main

import (
	"errors"
	"strings"
	"fmt"
)

// The general format of a function declaration starts with
// the keyword func, followed by the name of the function,
// then the list of parameters (data passed into the function)
// which is surrounded by parenthesis, then the return value(s)
// of the function which is also surrounded by parenthesis. IF
// and ONLY IF there is one return value, the parenthesis surrounding
// the return value can be omitted. If there are no return values, the
// entire return section of the function declaration can be
// omitted. After all of that is the function body which is
// surrounded by curly braces.
func addTwoNumbers(first, second int) int {
	return first + second
}

// For this next function, I will demonstrate what normal commenting
// for a function would look like, which appears in automatically
// generated godocs if your repository is hosted on GitHub.

// Function appendToString takes parameters for unappended string,
// a string to append to the unappended string, and the amount of
// times to append the string to the unappended string. If the given
// strings are empty or the count is less than 1 the function will
// return an error.
func appendToString(unappendedString, stringToAppend string, count int) (string, error) {
	if len(strings.TrimSpace(unappendedString)) == 0 || len(strings.TrimSpace(stringToAppend)) == 0 {
		return "", errors.New("given strings cannot be empty")
	}

	if count <= 0 {
		return "", errors.New("count cannot be less than 1")
	}

	for i := 0; i < count; i++ {
		unappendedString += stringToAppend
	}

	return unappendedString, nil
}

func main() {
	sum := addTwoNumbers(1, 2)
	fmt.Println(sum)

	// To capture multiple return values from a function, the
	// values variables are just separated by commas.
	newString, err := appendToString("Join the Gophers Slack", "!", sum)
	if err != nil {
		fmt.Printf("An error has occurred: %v\n", err.Error())
		return
	}
	fmt.Println(newString)

	// To ignore select return values from multiple return
	// functions the underscore operator must be used.
	anotherNewString, _ := appendToString("I know this won't generate an error.. or at least I hope it won't", "!", 1)
	fmt.Println(anotherNewString)
}
