package main

import "fmt"

type duration int

// There are some errors in the following
func main() {
	var hour duration = 3600
	minute := 60
	// fmt.Println(hour != minute)
	fmt.Println(hour != duration(minute))
}
