package main

// Recursion is a function that calls itself
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	println(factorial(5))
}
