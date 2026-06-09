package main

import "fmt"

func main() {
	// Consider the following string declaration:
	s1 := "țară means country in Romanian"

	// 1. Iterate over the string and print out byte by byte
	for i := range s1 {
		fmt.Printf("%b\n", s1[i])
	}

	fmt.Println()

	// 2. Iterate  over the string and print out rune by rune and the byte index where the rune starts in the string
	for i, c := range s1 {
		fmt.Printf("%d %d\n", i, c)
	}
}
