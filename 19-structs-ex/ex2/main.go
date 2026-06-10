package main

import "fmt"

func main() {
	// Consider the code from the previous exercise and:
	type person struct {
		name   string
		age    int
		colors []string
	}

	me := person{name: "Mario", age: 51, colors: []string{"black", "blue"}}
	you := person{name: "Mariarosa", age: 50, colors: []string{"green"}}

	// 1. Change the name or the struct value called me to "Andrei"
	me.name = "Andrei"

	// 2. Take in a new variable the favorites colors of struct value called you.
	// Print out the type and the value of the new variable.
	yourColors := you.colors
	fmt.Printf("Your colors: %#v\n", yourColors)

	// 3. Add a new favorite color to the second struct value called you.
	yourColors = append(yourColors, "pink")
	you.colors = yourColors

	// 4. Print out the struct values.
	fmt.Printf("Your colors: %#v\n", you.colors)
}
