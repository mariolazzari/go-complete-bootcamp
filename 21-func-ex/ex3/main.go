package main

import (
	"fmt"
	"log"
	"strconv"
)

// Write a function called myFunc() that takes exactly one argument which is an int number written between double quotes
// (this is in fact a string).
//
// If the argument is integer 'n', the function should return the result of n + nn + nnn
// Example: myFunc('5') returns 5 + 55 + 555 which is 615 and myFunc('9') returns 9 + 99 + 999 which is 1107

func myFunc(n string) int {
	num, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	}
	return num * (1 + 11 + 111)
}

func main() {
	fmt.Println(myFunc("5"))
	fmt.Println(myFunc("9"))
}
