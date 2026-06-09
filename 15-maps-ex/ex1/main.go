package main

import "fmt"

func main() {
	// 1. Declare a map called m1 which value is nil. Print out its type and value.
	var m1 map[int]string
	fmt.Printf("%#v\n", m1)

	// 2. Declare a map called m2. It's keys are of type int and values of type string.
	// Initialize the map using  a map literal with two key:value pairs.
	m2 := map[int]string{1: "Mario"}
	fmt.Printf("%#v\n", m2)

	// 3. Add the following key: value to the map: 10: "Abba"
	m2[10] = "Abba"

	// 4. Retrieve the value of an existing key and the value of a non existing key. Think about the results.
	val, ok := m2[1]
	if ok {
		fmt.Printf("Val 1 exists: %s\n", val)
	} else {
		fmt.Println("Val 1 does not exist")
	}

	val, ok = m2[11]
	if ok {
		fmt.Printf("Val 11 exists: %s\n", val)
	} else {
		fmt.Println("Val 11 does not exist")
	}
}
