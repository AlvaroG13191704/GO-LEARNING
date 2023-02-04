package main

// Pointers are variables that store the memory address of another variable

// We pass the address of a variable to a function, and the function can modify the value of the variable no matter where it's actually defined

func increment(inc *int) {
	*inc++ // dereference the pointer and increment the value at the address
}

func main() {
	x := 0
	println("x before", x) // x is 0

	increment(&x)         // pass the memory address of x
	println("x after", x) // x is 1
}
