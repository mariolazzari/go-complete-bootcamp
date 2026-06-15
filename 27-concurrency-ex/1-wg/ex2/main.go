package main

import (
	"fmt"
	"sync"
)

// 1. Create a function called sum() that calculates and then prints out  the sum of 2 float numbers it receives as arguments.
// Format the result with 2 decimal points.
func sum(a, b float64, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
}

func main() {

	// 2. From main launch 3 goroutines that execute the function you have just created (sum)
	// 3. Synchronize the goroutines and the main function using WaitGroups

	var wg sync.WaitGroup

	wg.Add(3)

	go sum(1, 2, &wg)
	go sum(4, 5, &wg)
	go sum(0, -1, &wg)

	wg.Wait()
}
