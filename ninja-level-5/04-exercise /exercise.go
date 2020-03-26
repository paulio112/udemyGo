package main

import "fmt"

func main() {
	anon := struct {
		name string
		age  int
	}{
		name: "Paul Southall",
		age:  32,
	}
	fmt.Println(anon)
}

// Hands-on exercise #4
// Create and use an anonymous struct.
// solution: https://play.golang.org/p/otBHFs-lPp
// video: 089
