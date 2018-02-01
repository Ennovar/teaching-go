package main

import "fmt"

func main() {
	num := 10

	// If, else, and else if statements. Note that the each subsequent condition
	// is on the same line as the preceding condition's ending curly-brace. This
	// is required, the compiler will throw an error if it isn't on the same line.
	if num > 5 {
		fmt.Printf("%v is greater than 5\n", num)
	} else if num < 2 {
		fmt.Printf("%v is less than 2\n", num)
	} else {
		fmt.Printf("%v is neither greater than 5 or less than 2\n", num)
	}

	// Typically, to make code easier to read, Go programmers try to avoid else
	// and else if conditions.
	var exampleString = "The number is less than 5"
	if num > 5 {
		exampleString = "The number is greater than 5"
	}
	fmt.Println(exampleString)

	// Switch statement syntax
	switch num {
	case 4:
		fmt.Println("The variable num holds a value of 4")
	case 7:
		fmt.Println("The variable num holds a value of 7")
	default:
		fmt.Println("This switch case doesn't have a case that can handle the variable num")
	}
}


