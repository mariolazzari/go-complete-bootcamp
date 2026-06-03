package main

import "fmt"

func main() {
	const days int = 7

	var i int
	fmt.Println(i)

	const pi = 3.14
	fmt.Println(days, pi)

	const a, b = 5, 0
	// fmt.Println(a/b) -> error detected
	fmt.Println(a, b)

	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)
	fmt.Println(min1, min2, min3)

	// group constant
	const (
		max1 = 100
		max2
		max3
	)
	fmt.Println(max1, max2, max3)

}
