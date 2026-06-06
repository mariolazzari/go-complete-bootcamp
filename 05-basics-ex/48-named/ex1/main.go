package main

import "fmt"

// Declare a new type type called duration. Have the underlying type be an int.
type duration int

// Declare a variable of the new type called hour using the var keyword
var hour duration

func main() {
	// print out the value of the variable hour
	fmt.Printf("%d\n", hour)
	// print out the type of the variable hour
	fmt.Printf("%T\n", hour)
	// assign 3600 to the variable hour using the  = operator
	hour = 3600
	// print out the value of hour
	fmt.Printf("%d\n", hour)
}
