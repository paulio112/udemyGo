package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:10])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}

// Took example from Udemy - Intentionally to make examples too.
// Using the code from the previous example, use SLICING to create the following new slices which are then printed:
// [42 43 44 45 46]
// [47 48 49 50 51]
// [44 45 46 47 48]
// [43 44 45 46 47]
// solution: https://play.golang.org/p/SGfiULXzAB
// video: 073
