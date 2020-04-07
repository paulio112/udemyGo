package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("JSON DIDNT LOAD")
	} else {

		// Note you wouldn't capture a success - Only done that for purpose of exercise
		fmt.Println("NO ERRORS WOOOP")
	}
	fmt.Println(string(bs))

}

// Hands-on exercise #1
// Start with this code. Instead of using the blank identifier, make sure the code is checking and handling the error.
// solution:
// https://play.golang.org/p/tn8oiuL1Yn
// video: 176
