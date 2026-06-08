package main

import "fmt"

func main() {
	mySlice := []float64{1.2, 5.6}

	// There are some errors
	mySlice[0] = 6

	// There are some errors
	a := float64(10)
	mySlice[0] = a

	// There are some errors
	// mySlice[3] = 10.10
	mySlice[1] = 10.10

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)
}
