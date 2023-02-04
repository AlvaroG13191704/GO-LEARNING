package main

import "strings"

// Closures functions are nested functions that can access variables from the scope in which they were defined

func repeat(n int) func(word string) string { // a func who receives an int and returns a func who receives a string and returns a string

	return func(word string) string {
		return strings.Repeat(word, n)

	}

}

func main() {
	repeat(5)("Hello") // HelloHelloHelloHelloHello

	// with variables
	repeat5 := repeat(5)
	repeat5("Hello") // HelloHelloHelloHelloHello
}
