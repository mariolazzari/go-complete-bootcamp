package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	const gr = 100
	// declaring a WaitGroup to synchronize the goroutines with the main function.
	var wg sync.WaitGroup

	// adding 200 goroutines to the WaitGroup
	wg.Add(gr * 2)

	// declaring a shared value
	var n int = 0

	// starting 200 goroutines
	for range gr {

		// each goroutine is an anonymous function
		go func() {
			// notifying the WaitGroup that the goroutine is done
			defer wg.Done()
			// introducing some artificial time
			time.Sleep(time.Second / 10)
			// increment the shared value
			n++
		}()

		// goroutine that decrements the shared value
		go func() {
			defer wg.Done()
			time.Sleep(time.Second / 10)
			n--
		}()

	}
	// waiting for the goroutines to terminate.
	wg.Wait()

	//  printing the final value of n
	fmt.Println(n) // it will have another value for each program execution -> DATA RACE
}
