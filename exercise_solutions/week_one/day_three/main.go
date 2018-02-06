package main

import "fmt"

type user struct {
	FirstName string
	LastName string
	Age int
}

func main(){
	usersInGroups := make(map[int][]user)

	// In an actual application, this data would be generated
	// by some sort of loop whilst most likely querying some
	// sort of database.
	usersInGroups[0] = []user{
		{"George", "Shaw", 18},
		{"John", "Doe", 22},
	}
	usersInGroups[1] = []user{
		{"Jane", "Doe", 16},
		{"Ashley", "Small", 43},
	}
	usersInGroups[2] = []user{
		{"Jimmy", "Parks", 24},
		{"Tom", "Smith", 65},
	}
	usersInGroups[3] = []user{
		{"Jessica", "Jones", 28},
	}

	// Loop through the map where k = the integer index and
	// v = the slice of users associated with that index
	for k, v := range usersInGroups {

		// Loop through v which is the current slice of users which
		// gives us the value user which is equal to the current user
		// within the slice
		for _, user := range v {
			fmt.Printf("%v %v is %v years old and is in group #%v\n", user.FirstName, user.LastName, user.Age, k)
		}
	}
}
