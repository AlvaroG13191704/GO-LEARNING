package main

import "fmt"

func main() {
	fmt.Println("Hello, World!") // To make line jump

	name := "John"
	fmt.Printf("Hello, %s \n", name) // Format string

	message := fmt.Sprintf("Hello, %s \n", name) // Format string and return string

	fmt.Println(message)

	fmt.Printf("This variable is ->  %T \n", message) // To know the type of variable use %T

	fmt.Print("Enter your age: ")
	var age int
	fmt.Scan(&age) // To get input from user
	fmt.Println("Your age is: ", age)
}
