package main

import "fmt"

func main() {
	// Consider the code from Exercise #1.
	type person struct {
		name   string
		age    int
		colors []string
	}

	me := person{name: "Mario", age: 51, colors: []string{"black", "blue"}}

	// Iterate and print out the favorite colors of the struct value called me.
	for _, color := range me.colors {
		fmt.Println(color)
	}
}
