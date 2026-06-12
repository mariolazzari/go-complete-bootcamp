package main

import "fmt"

func main() {
	// Consider the following variable declarations:
	x, y := 10, 2
	ptrx, ptry := &x, &y

	// Declare a new variable called z and initialize the variable by dividing x by y through the pointers.
	z := *ptrx / *ptry
	fmt.Printf("z = %v", z)
}
