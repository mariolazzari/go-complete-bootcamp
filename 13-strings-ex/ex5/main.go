package main

import "fmt"

func main() {
	// Consider the following string declaration:
	s := "你好 Go!"

	// 1. Convert the string to a byte slice.
	b := []byte(s)

	// 2. Print out the byte slice
	fmt.Printf("%b\n", b)

	// 3. Iterate over the byte slice and print out each index and byte in the byte slice
	for i, v := range b {
		fmt.Printf("%d %b\n", i, v)
	}
}
