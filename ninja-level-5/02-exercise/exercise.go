package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	favFlavors []string
}

func main() {

	p1 := person{
		first_name: "Paul",
		last_name:  "Southall",
		favFlavors: []string{
			"mint",
			"tea",
			"Other",
		},
	}
	p2 := person{
		first_name: "Claire",
		last_name:  "Lake",
		favFlavors: []string{
			"Salted Caramel",
			"Choc",
			"Strawberry",
		},
	}

	//Sample how to map from a construct based on a key of last name.
	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for _, v := range m {
		fmt.Println(v.first_name)
		fmt.Println(v.last_name)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}

// Subsequent build on top of :
// Take the code from the previous exercise, then store the values of type person in a map
// with the key of last name. Access each value in the map.
// Print out the values, ranging over the slice.
// solution: https://play.golang.org/p/RvvLyAMvGo
// video: 087

//  Exercise info

// Create your own type “person” which will have an underlying type of “struct” so
// that it can store the following data:
// first name
// last name
// favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements
// in the slice which stores the favorite flavors.
// solution:
// https://play.golang.org/p/ouRHmH_POg
// video: 086
