package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 4):
		fmt.Println("This should not print")
	case (3 == 3):
		fmt.Println("This should print")
		fallthrough
	case (4 == 4):
		fmt.Println("This should print")
	case (5 == 5):
		fmt.Println("This should not print")
	}

	// switch with variable
	n := "Bond"
	switch n {
	case "Moneypenny", "Bond", "Q":
		fmt.Println("miss money or bond or Q")
	case "M":
		fmt.Println("m")
	default:
		fmt.Println("this is default")
	}
}
