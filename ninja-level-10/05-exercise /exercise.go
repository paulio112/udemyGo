package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 32
	}()
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

// Show the comma ok idiom starting with this code.
// solution: https://play.golang.org/p/qh2ywLB5OG
// video: 168
