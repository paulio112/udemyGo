package main

import "fmt"

func main() {

	favsport := "Tennis"
	switch favsport {
	case "tennis":
		fmt.Println("This is tennis ")
	case "soccer":
		fmt.Println("soccer")
	default:
		fmt.Println("It didn't match our switch statement")
	}
}

// Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
// solution: https://play.golang.org/p/h-8FnjpzWB
// video: 058
