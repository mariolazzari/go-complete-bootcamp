package main

import "fmt"

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	var x any
	x = cube{edge: 5}
	// There is an error in the following
	// v := volume(x)

	c, ok := x.(cube)
	if ok {
		fmt.Printf("Cube Volume: %.2f\n", volume(c))
	}
}
