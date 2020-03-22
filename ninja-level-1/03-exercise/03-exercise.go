package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	x = 43
	y = "james bond"
	z = true
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}

// Using the code from the previous exercise,
// At the package level scope, assign the following values to the three variables
// for x assign 42
// for y assign “James Bond”
// for z assign true
// in func main
// use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
// print out the value stored by variable “s”
// code: here’s the solution: https://play.golang.org/p/QFctSQB_h3
// video: 019
