package main

import "fmt"

func main() {
	myMap := make(map[int][]string)

	myMap[0] = []string{
		"Some data", "Other data", "Some more data",
	}
	myMap[1] = []string {
		"Here is more data", "This is random data",
	}
	myMap[2] = []string {
		"Data is everywhere", "There is so much random data",
	}

	for k, v := range myMap {
		fmt.Printf("Current key: %v\n", k)
		for _, innerV := range v {
			fmt.Printf("\tInner value: %v\n", innerV)
		}
		fmt.Print("\n");
	}
}
