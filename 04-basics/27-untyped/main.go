package main

import "fmt"

func main() {
	// typed constant
	const a float64 = 7.7
	// untyped constant
	const b = 6.6

	const c = a * b
	const name = "Mario" + "Lazzari"

	fmt.Println(c, name)

	// const x int = 5
	// const y float64 = 7 * x -> error: type mismatched

	const x = 5
	const y = x * 2.2

	fmt.Printf("%T\n", y)

	const i int = x     // untyped x changes to int
	const j float64 = x // untyped x changes to float
	fmt.Printf("%T %T\n", i, j)

}
