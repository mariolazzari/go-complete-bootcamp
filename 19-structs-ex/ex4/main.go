package main

import "fmt"

func main() {
	// Change the code from Exercise #1 and:

	// 1. Create a new struct type called grades with  2 fields: grade and course
	type grades struct {
		grade  int
		course string
	}

	// 2. Add another field of type grades to person struct type (embedded struct).
	type person struct {
		name   string
		age    int
		colors []string
		grades grades
	}

	// 3. Change the initialization of the struct values called me and you to contain information about the grades.
	me := person{name: "Mario", age: 51, colors: []string{"black", "blue"}, grades: grades{grade: 100, course: "a"}}

	// 4. Change the grade and the course of one struct value to "Golang" and 98.
	me.grades.course = "Golang"
	me.grades.grade = 98

	// 5. Print out the struct values.
	fmt.Printf("me: %#v\n", me)
}
