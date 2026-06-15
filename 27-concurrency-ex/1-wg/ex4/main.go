package main

import (
	"fmt"
	"math"
	"sync"
)

// Change the code from Exercise #3 and launch 50 goroutines that calculate concurrently
// the square root of all the numbers between 100 and 149 (both included).

func main() {
	var wg sync.WaitGroup
	wg.Add(50)

	for i := 100; i <= 149; i++ {
		go func(n int, wg *sync.WaitGroup) {
			defer wg.Done()
			x := math.Sqrt(float64(n))
			fmt.Printf("Square root of %d is %.4f\n", n, x)
		}(i, &wg)

	}
	wg.Wait()
}
