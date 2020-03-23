package main

import "fmt"

func main() {
	//
	x := [5]int{2, 3, 5, 7, 13}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}

// Using a COMPOSITE LITERAL:
// create an ARRAY which holds 5 VALUES of TYPE int
// assign VALUES to each index position.
// Range over the array and print the values out.
// Using format printing
// print out the TYPE of the array
// solution: https://play.golang.org/p/tD0SzV3hdf
// video: 071
