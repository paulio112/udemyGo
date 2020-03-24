package main

import "fmt"

func main() {

	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
		`to_remove`:       []string{`zero`, `one`, `two`},
	}

	//Line to add to the map below Map[key] = []type{data}
	m["paul_southall"] = []string{"SLT", "gant charts", "meetings"}

	// Remove example .
	delete(m, "to_remove")

	for k, val := range m {
		fmt.Println("information for", k)
		for i, val2 := range val {
			fmt.Println("\t", i, val2)
		}
	}
}

// Misc data for exercise 10
// Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
// solution: https://play.golang.org/p/TYl5EbjoeC
// video: 080

// Using the code from the previous example, add a record to your map.
//Now print the map out using the “range” loop
// solution: https://play.golang.org/p/_CkxAhRrDJ
// video: 079

// Create a map with a key of TYPE string which is a person’s “last_first” name, and a
//  value of TYPE []string which stores their favorite things.
//   Store three records in your map. Print out all of the values,
//    along with their index position in the slice.

// 	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
// `moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
// `no_dr`, `Being evil`, `Ice cream`, `Sunsets`

// solution: https://play.golang.org/p/nTzSlRa9_A
// video: 078
