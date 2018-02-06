package main

import "fmt"

// User is now a special data type that can be utilized
// just like any other data type in Go now that it is
// declared.
type User struct {
	FirstName string
	LastName string
	Age int32
}

// Struct types can implement functions attached to them
func (u *User) PrettyPrint() {
	fmt.Printf("Hello, my name is %v %v. I am %v years old!\n", u.FirstName, u.LastName, u.Age)
}

func main() {
	// Here is one way to assign values to a struct
	me := User{
		FirstName: "George",
		LastName: "Shaw",
		Age: 18,
	}
	me.PrettyPrint()

	// Here is another way to assign values to a struct
	john := User{"John", "Doe", 54}
	john.PrettyPrint()

	// Here is how to reassign values within a struct
	john.LastName = "Smith"

	// Empty structs still contain values, the values are
	// just the zero-values of the variable types
	var emptyUser User
	emptyUser.PrettyPrint()
}