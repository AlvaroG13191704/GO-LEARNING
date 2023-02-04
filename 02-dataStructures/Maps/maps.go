package main

import "fmt"

// Maps are Go's built-in associative data type (sometimes called hashes or dicts in other languages).
// Maps cab be complex colletions of data, but they are not the same as arrays or slices.

func main() {
	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	days := make(map[int]string)
	fmt.Println(days)
	// Set key/value pairs using typical name[key] = val syntax.
	// add values
	days[1] = "Monday"
	days[2] = "Tuesday"
	days[3] = "Wednesday"
	days[4] = "Thursday"
	days[5] = "Friday"
	days[6] = "Saturday"
	fmt.Println(days)

	// Add values
	days[7] = "Sunday"
	fmt.Println(days)

	// Delete values
	delete(days, 7)

	// New map with arrays
	students := make(map[string][]int)
	students["John"] = []int{1, 2, 3, 4, 5}
	students["Mary"] = []int{1, 2, 3, 4, 5}
	students["Peter"] = []int{1, 2, 3, 4, 5}
	fmt.Println(students)

	// access values
	fmt.Println(days[1])
	fmt.Println(students["John"][2])
}
