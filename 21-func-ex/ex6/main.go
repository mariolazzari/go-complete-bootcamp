package main

import (
	"fmt"
	"strings"
)

// Create a function called searchItem() that takes 2 parameters:
// a) a string slice and
// b) a string.

// The function should search for the string (the second parameter) in the slice
// (the first parameter) and returns true if it finds the string in the slice and false otherwise.
// The function does a case-sensitive search.

func searchItem(vals []string, search string) bool {
	for _, val := range vals {
		if strings.Contains(val, search) {
			return true
		}
	}
	return false
}

func main() {
	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "pig")
	fmt.Println(result) // => false
}
