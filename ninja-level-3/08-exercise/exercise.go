package main

import "fmt"

func main() {
	i := 3
	switch i {
	case 1:
		fmt.Println("I is 1")
	case 2:
		fmt.Println("I is 2")
	default:
		fmt.Println("I didn't match our switch statement")
	}

}

// Create a program that uses a switch statement with no switch expression specified.
// solution: https://play.golang.org/p/YpPgkWGqKY
// video: 057
