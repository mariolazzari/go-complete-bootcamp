package main

import "fmt"

// Modify only the line in the main() body function where the print() function is invoked
// so that the program will print out Hello, Go playground! and then
// The Go gopher is the iconic mascot of the Go project.

func print(msg string) {
	fmt.Println(msg)
}

func main() {
	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")
}
