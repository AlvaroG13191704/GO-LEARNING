package main

import "fmt"

func main() {

	// If statement similar to C
	a := 10
	b := 20
	if a < b {
		fmt.Println("a < b")
	}

	// If-else statement
	if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a <= b")
	}

	// Nested if statement
	if a < b {
		if a > 0 {
			fmt.Println("a < b && a > 0")
		}
	}

	// If as a variable
	// This is a short declaration statement
	// It declares a variable and assigns a value to it
	// It can only be used inside a function
	if r := a < b; r {
		fmt.Println("a < b")
	}

	// Mini example
	users := make(map[string]string)
	users["Alvaro"] = "ga@gmail.com"
	users["John"] = "joh@gmail.com"
	users["Jane"] = "jane@gmail.com"

	if email, ok := users["Alvaro"]; ok {
		fmt.Println("Alvaro's email is", email)
	} else {
		fmt.Println("Alvaro's email not found")
	}

}
