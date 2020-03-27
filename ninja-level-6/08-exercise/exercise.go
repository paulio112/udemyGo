package main

import "fmt"

func main() {
	//Assign the function foo to a variable.
	v := foo()
	// Return the variable v which has a function assigned to tupe.
	fmt.Println(v())
}

// First line a function that returns a function.
func foo() func() int {
	return func() int {
		return 12
	}
}

// Hands-on exercise #8
// Create a func which returns a func
// assign the returned func to a variable
// call the returned func
// code: https://play.golang.org/p/c2HwqVE1Rd
// video: 109
