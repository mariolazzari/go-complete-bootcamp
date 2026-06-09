package main

import "fmt"

func main() {
	// Consider the following string declaration:
	s := "你好 Go!"

	// 1. Convert the string to a rune slice.
	r := []rune(s)

	// 2. Print out the rune slice
	fmt.Printf("%v\n", r)

	// 3. Iterate over the rune slice and print out each index and rune in the rune slice
	for i, v := range r {
		fmt.Printf("%d %c\n", i, v)
	}
}
