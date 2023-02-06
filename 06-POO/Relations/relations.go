package main

// Relations one to one
type Person struct {
	Name  string
	Age   int
	Email string
}
type Employee struct {
	employee Person // this is the relation, different from the other example of inheritance
	Job      string
	Salary   int
}

// Relation one to many
type Course struct {
	Name  string
	Video []Video // A course can have many videos
}
type Video struct {
	Name   string
	Course Course // A video can only have one course
}

// Relations one to many
func main() {
	// one to one
	person := Person{Name: "John", Age: 30, Email: "adfas@gmail.com"}
	employee := Employee{employee: person, Job: "Developer", Salary: 1000} // an employee can only have one person and a person can only have one employee
	println(employee.employee.Name)

	// one to many
	video1 := Video{Name: "Video 1"}
	video2 := Video{Name: "Video 2"}
	video3 := Video{Name: "Video 3"}
	course := Course{Name: "Course 1", Video: []Video{video1, video2, video3}}
	println(course.Video[0].Name)
}
