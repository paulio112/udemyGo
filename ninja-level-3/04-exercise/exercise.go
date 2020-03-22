package main

import "fmt"

func main() {

	i := 1985
	for {
		if i > 2017 {
			break
		}
		fmt.Println(i)
		i++
	}
}

// Create a for loop using this syntax
// for { }
// Have it print out the years you have been alive.
// solution: https://play.golang.org/p/9VpnB-I1Pz
// video: 053
