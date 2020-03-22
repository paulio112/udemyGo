package main

import "fmt"

func short() {
	x := 1
	fmt.Println(x)
	x = 3
	fmt.Println(x)

	if x < 4 {
		fmt.Println("huge")
	} else {
		fmt.Println("not huge")
	}

}
