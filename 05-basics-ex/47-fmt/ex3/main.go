package main

import "fmt"

// There are some errors i
func main() {
	x, y := 4, 5.1

	// z := x * y
	z := float64(x) * y
	fmt.Println(z)

	a := 5
	// b := 6.2 * a
	b := 6.2 * float64(a)
	fmt.Println(b)
}
