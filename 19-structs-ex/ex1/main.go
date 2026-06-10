package main

import "fmt"

func main() {
	// 1. Create your own struct type called person that stores the following data:
	// name, age and a list with favorite colors.
	type person struct {
		name   string
		age    int
		colors []string
	}

	// 2. Declare and initialize two values of type person, one called me and another called you.
	me := person{name: "Mario", age: 51, colors: []string{"black", "blue"}}
	you := person{name: "Mariarosa", age: 50, colors: []string{"green"}}

	// 3. Print out the struct values.
	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", you)
}
