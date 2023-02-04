package main

import "fmt"

func main() {
	a := 10
	b := 20
	var r bool
	// == equal
	r = a == b
	fmt.Println("a == b", r)

	// != not equal
	r = a != b
	fmt.Println("a != b", r)

	// < less than
	r = a < b
	fmt.Println("a < b", r)

	// > greater than
	r = a > b
	fmt.Println("a > b", r)

	// <= less than or equal
	r = a <= b
	fmt.Println("a <= b", r)

	// >= greater than or equal
	r = a >= b
	fmt.Println("a >= b", r)

	// && and
	r = a < b && a > 0
	fmt.Println("a < b && a > 0", r)

	// || or
	r = a < b || a > 0
	fmt.Println("a < b || a > 0", r)

	// ! not
	r = !(a < b)
	fmt.Println("!(a < b)", r)

}
