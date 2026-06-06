package main

import "math"

// Try to identify the errors,
func main() {
	const a int = 7
	const b float64 = 5.6
	// const c = a * b
	const c = float64(a) * b

	// x := 8
	const x = 8
	const xc int = x

	// const noIPv6 = math.Pow(2, 128)
	noIPv6 := math.Pow(2, 128)
	_ = noIPv6
}
