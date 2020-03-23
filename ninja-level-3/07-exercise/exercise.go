package main

import "fmt"

func main() {

	for i := 10; i < 100; i++ {
		if i == 11 {
			fmt.Println("Int is 11")
		} else if i == 12 {
			fmt.Println("Int is 12")
		} else {
			fmt.Println("number is not 11 or 12")
		}
	}
}

// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
// solution: https://play.golang.org/p/IDnrJpE7vT
// video: 056
