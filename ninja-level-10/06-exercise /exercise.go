package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

// write a program that
// puts 100 numbers to a channel
// pull the numbers off the channel and print them
// solution: https://play.golang.org/p/096Lk1BR7o
