package main

import "fmt"

// Create a function that takes in an int value and prints out that value.
// Assign the function to a variable, print out the type of the variable and then call that function through the variable name.

func printN(n int) {
	fmt.Println(n)
}

func main() {
	p := printN
	fmt.Printf("%T\n", p)
	p(1)
}
