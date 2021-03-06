package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello, playground")
}

func foo() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	defer func() {
		fmt.Println("foo DEFER ran2")
	}()
	fmt.Println("foo ran")
}

// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
// code: https://play.golang.org/p/XluEuUD0Nw
// video: 104
