package main

import "fmt"

// Write a Go program that converts i to float64 and f to int.
func main() {
	var i int = 3
	var f float64 = 3.2

	iF := float64(i)
	fI := int(f)

	// Print out the type and the value of the variables created after conversion.
	fmt.Printf("i type: %T, f type: %T\n", iF, fI)
}
