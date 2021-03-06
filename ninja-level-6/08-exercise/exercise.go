package main

import "fmt"

func main() {
	a := foo()
	fmt.Println(a())
}

// foo is a function that accpets a function.

func foo() func() int {
	return func() int {
		return 112
	}
}

// Hands-on exercise #8
// Create a func which returns a func
// assign the returned func to a variable
// call the returned func
// code: https://play.golang.org/p/c2HwqVE1Rd
// video: 109
