package main

import (
	"fmt"
	"strings"
)

// Change the function from the previous exercise to do a case-insensitive search.

func searchItem(vals []string, search string) bool {
	for _, val := range vals {
		if strings.EqualFold(val, search) {
			return true
		}
	}
	return false
}

func main() {
	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "Bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "Pig")
	fmt.Println(result) // => false
}
