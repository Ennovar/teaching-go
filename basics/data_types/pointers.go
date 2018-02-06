package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Declare and initialize myInt to the value 100
	myInt := 100

	// Declare and initialize the variable pointer to the memory
	// address of the variable myInt
	pointing := &myInt

	// Pointers don't have a special type in Go, they are just
	// the default type of whatever variable it points to with
	// an asterisk before the type to denote that it is a pointer
	// to that type
	fmt.Printf("Variable \"pointing\" is of type: %v\n", reflect.TypeOf(pointing))

	// Read the value of the variable myInt through the memory
	// address stored in the variable pointing by using the
	// asterisk operator which references the value within the
	// memory address
	fmt.Println(*pointing)

	// By using the asterisk operator to reference the underlying
	// value of the pointer, I can reassign the variable first to
	// the value of 150. The value of pointer stays the same, because
	// it is a memory address, the only thing that changes is the
	// value within that memory address
	*pointing = 150
	fmt.Println(*pointing)

	// By using the asterisks operator again to reference the value
	// of the memory address stored in the integer pointer, pointing,
	// the value of myInt can be reassigned to the value of myInt plus
	// 50.
	*pointing = *pointing + 50
	fmt.Println(*pointing)
}
