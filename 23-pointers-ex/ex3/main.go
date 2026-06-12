package main

import "fmt"

func main() {
	// Consider the following variable declaration:
	x, y := 5.5, 8.8

	// Write a function that swaps the values of x and y.
	// After calling the function x will be 8.8 and y will 5.5

	fmt.Printf("x = %f, y = %f\n", x, y)
	swap(&x, &y)
	fmt.Printf("x = %f, y = %f\n", x, y)
}

func swap(a, b *float64) {
	*a, *b = *b, *a
}
