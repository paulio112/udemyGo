package main

import "fmt"

// Can also use constants as per below.
// const {
// 	a = 3
// 	b = 4
// }

func main() {
	const a = 3
	const b int = 4
	fmt.Println(a, b)
}

// Create TYPED and UNTYPED constants. Print the values of the constants.
// solution: https://play.golang.org/p/NutvJXWUx2
// video: 033
