package main

import "fmt"

type Animal interface { // interface is a type that can have methods but not attributes
	move() string
}

type Dog struct { // Dog and Fish are the implementations of the interface Animal
}

type Fish struct {
}

func (d *Dog) move() string { // the methods of the interface must be implemented in the struct
	return "Dog is moving"
}

func (f *Fish) move() string {
	return "Fish is moving"
}

func moveAnimal(animal Animal) { // This function receives an interface as a parameter
	fmt.Println(animal.move())
}

func main() {
	dog := Dog{}
	moveAnimal(&dog) // the interface is passed by reference

	fish := Fish{}
	moveAnimal(&fish)
}
