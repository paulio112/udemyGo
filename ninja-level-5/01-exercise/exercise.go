package main

import "fmt"

type person struct {
	first_name   string
	last_name    string
	fav_icecream []string
}

func main() {

	p1 := person{
		first_name: "Paul",
		last_name:  "Southall",
		fav_icecream: []string{
			"choc",
			"tea",
			"Other",
		},
	}
	p2 := person{
		first_name: "Claire",
		last_name:  "Lake",
		fav_icecream: []string{
			"orange",
			"mint",
			"banana",
		},
	}

	// using range to go over ice cream data for individual person.
	fmt.Println(p1.first_name)
	fmt.Println(p1.last_name)

	for i, v := range p1.fav_icecream {
		fmt.Println(i, v)
	}

	// Using to identify p2
	fmt.Println(p1.first_name)
	fmt.Println(p1.last_name)

	for i, v := range p2.fav_icecream {
		fmt.Println(i, v)
	}
}

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
