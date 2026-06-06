package main

import (
	"fmt"
	"strconv"
)

// Write a Go program that converts
// // Print the value and the type for each variable created after conversion.
func main() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	// i to string (int to string)
	s := strconv.Itoa(i)
	fmt.Printf("s Type is %T, s value is %q\n", s, s)

	// s2 to int (string to int)
	is, err := strconv.Atoi(s2)
	if err == nil {
		fmt.Printf("i type is %T, i value is %v\n", is, is)
	} else {
		fmt.Println("Can not convert string to int.")
	}

	// f to string (float64 to string)
	ss1 := fmt.Sprintf("%f", f)
	fmt.Printf("ss1's type: %T, ss1's value: %s\n", ss1, ss1)

	// s1 to float64  (string to float64)
	f1, err1 := strconv.ParseFloat(s1, 64)
	if err1 == nil {
		fmt.Printf("f1's type: %T, f1's value: %v\n", f1, f1)
	} else {
		fmt.Println("Cannot convert string to float64.")
	}
}
