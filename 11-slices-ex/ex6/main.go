package main

import "fmt"

func main() {
	// Consider the following slice declaration:
	friends := []string{"Marry", "John", "Paul", "Diana"}

	// Using copy() function create a copy of the slice.
	friendsCopy := make([]string, len(friends))
	_ = copy(friendsCopy, friends)

	// Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.
	friends[0] = "Mario"
	fmt.Printf("original: %#v\n", friends)
	fmt.Printf("copy: %#v\n", friendsCopy)
}
