package main

import "fmt"

func main() {

	// // Example one is to remove the specificness from the channel.
	// cs := make(chan int)

	// go func() {
	// 	cs <- 42
	// }()
	// fmt.Println(<-cs)

	// fmt.Printf("------\n")
	// fmt.Printf("cs\t%T\n", cs)

	//
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

}

// Hands-on exercise #4
// Starting with this code, pull the values off the channel using a select statement
// solution: https://play.golang.org/p/FulKBY5JNj
// video: 167
