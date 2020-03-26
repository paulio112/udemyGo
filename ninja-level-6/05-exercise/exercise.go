package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

//expects type of function: - circle,  returns an area which is of type float64
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//Expects a type of function - square  and returns the area of the function
func (s square) area() float64 {
	//To calculate the area of a square it's lenth  *  length
	return s.length * s.length
}

// expects a shape param, this allow data available via interface .
func info(s shape) {
	fmt.Println(s.area())
}

func main() {

	circ := circle{
		radius: 50,
	}

	squa := square{
		length: 10,
	}

	//calling a method of info with TYPE circle assigned to a variable
	info(circ)
	info(squa)

}

// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// circle area= π r 2
// square area = L * W
//
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle
// code: https://play.golang.org/p/NGGikHNvHv
// video: 106
