package main

import "fmt"

func sum(nums ...int) { // like args in python, we can pass any number of variables to this function

	var result int

	fmt.Println(nums) // this will print a slice of int

	for _, value := range nums {
		result += value
	}
	fmt.Println(result)
}

// Return multple values from a function
func sub(name string, nums ...int) (string, int) { // we can return multiple values from a function
	return name, nums[1]

}

func main() {
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	name, values := sub("Raj", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(name, values)
}
