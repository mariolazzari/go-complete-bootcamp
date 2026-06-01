package main

import (
	"fmt"
	"math"
)

func main() {
	x := 42.25
	res := math.Pow(x, 8)
	fmt.Printf("%.2f at pawer of 8: %.2f\n", x, res)
}
