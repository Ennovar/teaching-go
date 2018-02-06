package main

import "fmt"

var myName string
var myInt uint8

func main(){
	myName = "George"
	myInt = 234
	myBool := false

	if myBool {
		fmt.Println("myBool is true!")
	} else {
		fmt.Println("myBool is false.")
	}

	switch myInt {
	case 200:
		fmt.Println("Its 200!")
	default:
		fmt.Println("Its not 200!")
	}
}
