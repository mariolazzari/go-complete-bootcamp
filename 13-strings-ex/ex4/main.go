package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s1 := "Go is cool!"
	x := s1[0]
	fmt.Println(x)

	// There are some errors
	// s1[0] = 'x'
	strings.Replace(s1, s1, "G", 'x')

	// printing the number of runes in the string
	fmt.Println(len(s1))
	fmt.Println(utf8.RuneCountInString(s1))
}
