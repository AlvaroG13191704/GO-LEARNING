package main

// Anonymous functions are functions without a name
// They are often used as arguments to other functions
// or as a way to create closures

func main() {

	// Closures are functions that can access variables from the scope in which they were defined
	func() {
		println("Hello World from a closure")
	}() // those () are used to call the function

	// function assigned to a variable
	myFunc := func() {
		println("Hello World from a variable")
	}

	myFunc() // if we only call the function without the () we will get the memory address of the function
	// fmt.Println(myFunc) // 0x10a0e60
}
