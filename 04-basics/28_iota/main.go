package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)

	const (
		c11 = iota
		c22
		c33
	)
	fmt.Println(c11, c22, c33)

	const (
		North = iota // default 0
		Sud
		East
		West
	)
	fmt.Println(West) // 3

	const (
		a = iota * 2
		b
		c
	)
	fmt.Println(a, b, c) // 0,2,4

	// skip values
	const (
		x = -(iota + 2)
		_
		y
		z
	)
	fmt.Println(x, y, z) // -2,-4,-5
}
