package main

import "fmt"

func main() {
	// Using the var keyword, declare an array called cities with 2 elements of type string. Don't initialize the array.
	var cities [2]string
	// Print out the cities array and notice the value of its elements.
	fmt.Println(cities)

	// Declare an array called grades of type [3]float64 and initialize only the first 2 elements using an array literal.
	var grades [3]float64 = [3]float64{10, 20}
	// Print out the grades array and notice the value of its elements.
	fmt.Println(grades)

	// Declare an array called taskDone using the ellipsis operator. The elements are of type bool. Print out taskDone.
	var taskDone = [...]bool{true, false, true}
	fmt.Println(taskDone)

	// Iterate over the array called cities using the classical for loop syntax and the len function and print out element by element.
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	// Iterate over grades using the range keyword and print out element by element.
	for _, grade := range grades {
		fmt.Println(grade)
	}
}
