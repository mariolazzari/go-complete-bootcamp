package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Create a goroutine named power() that has one parameter of type int,
// calculates the square value of its parameter
// and then sends  the result into a channel.
func power(n int, ch chan<- int) {
	ch <- n * n
}

func main() {
	const jobs = 50
	ch := make(chan int, jobs)

	var wg sync.WaitGroup
	wg.Add(jobs)

	// In the main function launch 50 goroutines that calculate
	// the square values of all numbers between 1 and 50 included.
	for i := 1; i <= jobs; i++ {
		go func(n int) {
			defer wg.Done()
			power(n, ch)
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	fmt.Println(runtime.NumGoroutine())

	// Print out the square values.
	for val := range ch {
		fmt.Println(val)
	}
}
