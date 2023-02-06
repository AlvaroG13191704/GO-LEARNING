package main

import "fmt"

// Create a person struct
type Person struct {
	name string
	age  int
}

// Create the method speak
func (p *Person) speak() { // We use pointers to receive the object address, not a copy of the object
	fmt.Println("Hello, my name is", p.name, "and I am", p.age, "years old")
}

// Method to change the name
func (p *Person) changeName(newName string) {
	p.name = newName
}

// Inheritance
type Student struct {
	Person // Inheritance
	grade  int
}

func main() {
	// Inialize a person
	person1 := Person{name: "John", age: 20} // An object as JSON
	fmt.Println(person1)
	// change values
	person1.name = "John Doe"
	person1.age = 30
	fmt.Println(person1)

	// call the method speak
	person1.speak()
	// Change the name
	person1.changeName("John Doe")

	// Work with inheritance
	student1 := Student{Person: Person{name: "John", age: 20}, grade: 10}
	fmt.Println(student1) // {{John 20} 10}
}
