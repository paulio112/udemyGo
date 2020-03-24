package main

import "fmt"

func main() {

	type vehicle struct {
		doors  int
		colour string
	}

	type truck struct {
		vehicle
		fourwheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	paulTruck := truck{
		vehicle: vehicle{
			doors:  12,
			colour: "RED",
		},
		fourwheel: true,
	}

	// Intentionally seperate lines to match instruction.
	fmt.Println("TRUCK DETAILS")
	fmt.Println(paulTruck.doors)
	fmt.Println(paulTruck.colour)
	fmt.Println(paulTruck.fourwheel)

	paulSedan := sedan{
		vehicle: vehicle{
			doors:  5,
			colour: "blue",
		},
		luxury: false,
	}

	fmt.Println("SEDAN DETAILS")
	fmt.Println(paulSedan.doors)
	fmt.Println(paulSedan.colour)
	fmt.Println(paulSedan.luxury)

	//fmt.Println("paulTruck details:  ", paulTruck)
	fmt.Println("paulTruck details:  ", paulSedan)
}

// Exercise information
//Create a new type: vehicle.
// The underlying type is a struct.
// The fields:
// doors
// color
// Create two new types: truck & sedan.
// The underlying type of each of these new types is a struct.
// Embed the “vehicle” type in both truck & sedan.
// Give truck the field “fourWheel” which will be set to bool.
// Give sedan the field “luxury” which will be set to bool. solution
//
// Using the vehicle, truck, and sedan structs:
// using a composite literal, create a value of type truck and
//assign values to the fields;
//
// using a composite literal, create a value of type sedan and
//assign values to the fields.
//
// Print out each of these values.
// Print out a single field from each of these values.
//https://play.golang.org/p/PrTtTv_vVO
