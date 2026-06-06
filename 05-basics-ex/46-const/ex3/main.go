package main

import "fmt"

// Calculate how many seconds are in a year.
func main() {
	// Declare secPerDay constant and initialize it to the number of seconds in a day
	const secPerDay = 60 * 60 * 24
	// Declare daysYear constant and initialize it to 365
	const daysYear = 365
	// Use fmt.Printf() to print out the total number of seconds in a year.
	fmt.Printf("%d\n", secPerDay*daysYear)
}
