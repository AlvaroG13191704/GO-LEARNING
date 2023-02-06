package main

import "GOLEARNING/08-EDD/structs"

type Student struct {
	Name string
	Age  int
}

func main() {
	student := Student{"John", 20}

	// Create a new list
	list := structs.DoubleLinkedList{}
	// add values
	list.AddFirst(student)
	list.AddFirst(Student{"Mary", 21})
	list.AddFirst(Student{"Peter", 22})
	list.AddFirst(Student{"Paul", 23})
	list.AddFirst(Student{"Mark", 24})
	list.AddLast(Student{"Luke", 25})

	// Print the list
	list.Print()
}
