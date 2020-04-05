package main

import "fmt"

func main() {

	// Running a goroutin to avoid a deadlock
	// c := make(chan int)

	// go func() {
	// 	c <- 42
	// }()

	// fmt.Println(<-c)

	// Buffer example
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)

}

// Hands-on exercise #1
// get this code working:
// using func literal, aka, anonymous self-executing func
// solution: https://play.golang.org/p/SHr3lpX4so
// a buffered channel
// solution: https://play.golang.org/p/Y0Hx6IZc3U
// video: 164
