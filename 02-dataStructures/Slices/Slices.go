package main

import "fmt"

func main() {
	// Slices are dynamic
	// Slices are mutable
	// Slices are references to arrays
	// Slices are like pointers to arrays

	numbers := []int{1, 2, 3, 4, 5} // slice of 5 integers
	fmt.Println(numbers, len(numbers))

	// add more elements
	numbers = append(numbers, 6, 7, 8, 9, 10) // Like a desctructive method
	fmt.Println(numbers, len(numbers))

	// Sub slices
	subNumbers := numbers[0:5] // from index o to index 5
	fmt.Println(subNumbers)

	// When we made a change in a sub slice, we are changing the original array, because the slice is a reference to father array
	moths := []string{
		"January",
		"February",
		"March",
		"April",
	}
	// Print the pointers, length and capacity
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(moths), cap(moths), moths)

	// add an element
	moths = append(moths, "May")
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(moths), cap(moths), moths)

	// The memory reference changed because the capacity is not enough

}
