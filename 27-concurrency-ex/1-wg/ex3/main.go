package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	//	1. Create an anonymous function that calculates and prints out the square root of a float value it receives as argument.
	go func(f float64, wg *sync.WaitGroup) {
		defer wg.Done()
		x := math.Sqrt(f)
		fmt.Printf("Square root of %.2f is %.4f\n", f, x)
	}(16.1, &wg)

	// 2. Launch the function as a goroutine and synchronize it with main using WaitGroups
	wg.Wait()
}
