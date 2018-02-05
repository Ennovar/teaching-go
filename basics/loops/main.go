package main

import "fmt"

// There is only one loop in Go, the for loop

func main(){
	x := 1
	myInts := []int{
		1,2,3,4,5,6,7,8,9,8,7,6,5,4,3,2,1,
	}

	// The for loop can take simple arguments
	for x < 10 {
		fmt.Println(x)
		x++
	}

	fmt.Println()

	// Here is the structure of what most other
	// languages classify as a for loop
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	// It can also take more complex arguments. The usage
	// of the keyword range within for loops is very common
	// for looping through slices and maps.
	for _, v := range myInts {
		fmt.Println(v)
	}

	fmt.Println()

	// An infinite for loop, breaking through the usage
	// of the keyword "break". To continue in a loop in
	// Go, the same keyword as most other languages is
	// also used, "continue".
	for {
		fmt.Println("This would be an infinite for loop if it wasn't for the break")
		break
	}
}