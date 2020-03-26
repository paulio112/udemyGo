package main

import "fmt"

func main() {
	seq := []int{1, 2, 3, 4, 5, 6}
	n := foo(seq...)

	fmt.Println(n)
	fmt.Println(bar(seq))
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

//create a func with the identifier foo that
// takes in a variadic parameter of type int
// pass in a value of type []int into your func (unfurl the []int)
// returns the sum of all values of type int passed in
//
// create a func with the identifier bar that
// takes in a parameter of type []int
// returns the sum of all values of type int passed in
// code: https://play.golang.org/p/B0yRxtBQPD
// video: 103
