package main

import "fmt"

func main() {
	// Consider the following variable declaration
	x := 10.10

	// 1. Print out the address of x
	fmt.Printf("%p\n", &x)

	// 2. Declare a pointer called ptr that stores the address of x.
	ptr := &x

	// 3. Print out the type and the value of ptr.
	fmt.Printf("%T %v\n", ptr, ptr)

	// 4. Print the address of the pointer and the value of x though the pointer (use the dereferencing operator).
	fmt.Printf("%T %v\n", ptr, *ptr)
}
