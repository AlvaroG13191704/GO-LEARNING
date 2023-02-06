package main

import "fmt"

// Global variables are declared outside of functions

var a int = 10

// Constants are declared outside of functions

func modify() {
	a = 20
	fmt.Println(a)
}

func deferFunc() {
	a = 30
	fmt.Println("Defer function")
}
func main() {
	fmt.Println(a)
	defer deferFunc() // Defer function is executed at the end of the main function
	modify()
	fmt.Println(a)
}
