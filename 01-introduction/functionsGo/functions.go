package main

import "fmt"

// Create a function
func hello(name string) {

	fmt.Printf("Hello, %s \n", name)
}

// Functions who return a value
func sum(a int, b int) int {
	return a + b
}

func main() {
	// Call the function
	hello("Alvaro")
	// Return a value
	fmt.Println(sum(2, 3))
}
