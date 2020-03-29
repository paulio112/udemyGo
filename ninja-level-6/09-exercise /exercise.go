package main

import "fmt"

func main() {
	sample := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}
	x := foo(sample, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)
}

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

// Hands-on exercise #9
// A “callback” is when we pass a func into a func as an argument. For this exercise,
// pass a func into a func as an argument
// code: https://play.golang.org/p/0yGYPKh1y7
// video: 110
