package main

import "fmt"

// Create a function with the identifier sum that takes in a variadic parameter of type int
// and returns the sum of all values of type int passed in.
func sum(n ...int) int {
	total := 0
	for _, num := range n {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	fmt.Println(sum())
}
