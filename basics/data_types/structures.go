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

// Struct types can implement methods.
func (u *User) PrettyPrint() {
	fmt.Printf("Hello, my name is %v %v. I am %v years old!\n", u.FirstName, u.LastName, u.Age)
}

func main() {
	me := User{
		FirstName: "George",
		LastName: "Shaw",
		Age: 18,
	}
	me.PrettyPrint()

	john := User{"John", "Doe", 54}
	john.PrettyPrint()

	var emptyUser User
	emptyUser.PrettyPrint()
}