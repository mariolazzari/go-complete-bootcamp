package main

import (
	"fmt"
)

func main() {
	// 1. Declare a rune called r that stores the non-ascii letter ă
	r := 'ă'

	// 2. Print out the type of r
	fmt.Printf("%v\n", r)
	fmt.Printf("%T\n", r)

	// 3. Declare 2 strings that contains the values ma and m
	str1, str2 := "ma", "m"

	// 4. Concatenate the strings and the rune in a new string
	// (the new string will hold the value mamă ).
	str3 := str1 + str2 + string(r)
	fmt.Println(str3)
}
