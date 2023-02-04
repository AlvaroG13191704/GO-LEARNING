package main

import "fmt"

func main() {
	// arrays are static
	var numbers [5]int // array of 5 integers, the default value is 0
	fmt.Println(numbers)

	// change the values
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	fmt.Println(numbers)

	// Arrays are inmutable

	// Define an array with values
	names := [3]string{"Alvaro", "July", "Luis"} // array of 3 strings
	fmt.Println(names)

	// Dynamic arrays
	colors := [...]string{
		"Red",
		"Green",
		"Blue",
		"Yellow",
	}
	fmt.Println(colors, len(colors))

	// Arrays defined with index
	people := [...]string{
		0: "Alvaro",
		1: "July",
		3: "Luis",
		5: "Maria",
	}
	fmt.Println(people, len(people)) // In this case we have 6 elements, but the index 2 and 4 are empty

	// Sub arrays
	subNames := names[0:2] // from index o to index 2
	fmt.Println(subNames)
}
