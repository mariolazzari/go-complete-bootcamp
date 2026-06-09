package main

import "fmt"

func main() {
	// 1. Using the var keyword declare a string called name and initialize it with your name.
	var name = "Mario"
	// 2. Using short declaration syntax declare a string called country and assign the country you are living in to the string variable.
	country := "Italia"

	// 3. Print the following string on multiple lines like this:
	// Your name: `here the value of name variable`
	// Country: `here the value of country variable`
	fmt.Printf("Your name: '%s'\nCountry: '%s'\n", name, country)

	// 4. Print out the following strings:
	// a) He says: "Hello"
	// b) C:\Users\a.txt
	fmt.Println(`He says: "Hello"
C:\Users\a.txt`)
}
