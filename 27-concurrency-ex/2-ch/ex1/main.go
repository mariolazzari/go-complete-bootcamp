package main

import "fmt"

func main() {
	// 1. Using the var keyword, declare a bidirectional unbuffered channel called c1 that works with values of type float64
	var c1 chan float64

	// 2. Using the make() built-in function declare and initialize a receive-only channel called c2 and a send-only channel called c3.
	// Both work with data of type rune.
	c2 := make(<-chan rune)
	c3 := make(chan<- rune)

	// 3. Declare a bidirectional buffered channel  called c4 with a capacity of 10 ints.
	c4 := make(chan int, 10)

	// 4. Print out the type of all the channels declared.
	fmt.Printf("%T\n", c1)
	fmt.Printf("%T\n", c2)
	fmt.Printf("%T\n", c3)
	fmt.Printf("%T\n", c4)

}
