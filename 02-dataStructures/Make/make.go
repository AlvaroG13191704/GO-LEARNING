package main

import "fmt"

func main() {
	// make() is used to create slices, maps, and channels only.
	// make() returns an initialized (not zeroed) value of type slice, map, or channel (not just the type).

	numbers := make([]int, 4) // make a slice of 0 integers with a capacity of 3
	fmt.Println(numbers, len(numbers), cap(numbers))

	// change the values
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	fmt.Println(numbers, len(numbers), cap(numbers))
}
