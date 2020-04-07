package main

import "fmt"

func main() {

	c := make(chan int)

	for a := 0; a <= 10; a++ {

		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for v := range c {
		fmt.Println(v, c)
	}

}

// /Hands-on exercise #7
// write a program that
// launches 10 goroutines
// each goroutine adds 10 numbers to a channel
// pull the numbers off the channel and print them
// solutions:
// https://play.golang.org/p/R-zqsKS03P
// https://play.golang.org/p/quWnlwzs2z
// student solutions:
// https://twitter.com/mannion
// https://gist.github.com/mannion007/3c8899913974c1027ef6f13ec37b2b3f
// video: 170
