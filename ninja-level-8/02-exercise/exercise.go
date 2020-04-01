package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	//fmt.Println(s)

	var people []person
	//fmt.Println(people)

	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}

	// Example of what people type people being populated as
	// fmt.Println(people)

	for i, person := range people {
		// I is the index of the loop
		fmt.Println(" I is the following->", i)
		// person.  Allows us to get attriibutes of type person
		fmt.Println("\t", person.First, person.Last, person.Age)
		// sayings is a slice [] referenced via saying
		for _, saying := range person.Sayings {
			fmt.Println("\t\t", saying)
		}
	}
}

// 	Starting with this code, unmarshal the JSON into a Go data structure.
//Hint: https://mholt.github.io/json-to-go/
// code:
// code setup (just fyi, not needed for exercise):
// https://play.golang.org/p/nWPP5Z2Q4e
// solution:
// https://play.golang.org/p/r8oSG8DIPR
// video: 126

// Approach to tackle
// Step 1 - create a struct (this can be done by using the json to go struct url  )
// Step 2 - variable people of type person
// Step 3 - Step 3 -
