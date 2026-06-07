package main

import "fmt"

func main() {
	// Consider the following array declaration:
	nums := [...]int{30, -1, -6, 90, -6}

	// Write a Go program that counts the number of positive even numbers in the array.
	count := 0
	for _, n := range nums {
		if n > 0 && n%2 == 0 {
			count++
		}
	}
	fmt.Printf("Positive even numbers: %d\n", count)
}
