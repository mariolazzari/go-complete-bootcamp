package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255
	x++ // overflow
	fmt.Println(x)

	// a := int8(255 + 1) -> compile time overflow

	var b int8 = 127
	fmt.Printf("%d\n", b+1)

	b = -127
	b--
	fmt.Printf("%d\n", b)

	f := float32(math.MaxFloat32)
	fmt.Println(f)
}
