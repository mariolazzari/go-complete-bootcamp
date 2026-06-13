package main

import "fmt"

func main() {
	// 1. Declare a variable called empty of type empty interface. Print out its type.
	var empty any
	fmt.Printf("%T\n", empty)

	// 2. Assign an int value to the variable called empty. Print out its type.
	empty = 42
	fmt.Printf("%T\n", empty)

	// 3. Assign a float64 value to empty. Print out its type.
	empty = 42.42
	fmt.Printf("%T\n", empty)

	// 4. Assign an int slice value to empty. Print out its type.
	empty = []int{1, 2, 3}
	fmt.Printf("%T\n", empty)

	// 5. Type Assertion
	empty = append(empty.([]int), 10)

	// 6. Print out its type.
	fmt.Printf("%v\n", empty)
}
