package main

import "fmt"

func main() {
	// Using a composite literal declare and initialize a slice of type string with 3 elements.
	names := []string{"Mario", "Maria", "Mariarosa"}
	// Iterate over the slice and print each element in the slice and its index.
	for i, name := range names {
		fmt.Printf("%d - %s\n", i, name)
	}
}
