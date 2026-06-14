package main

import "fmt"

func factorial(n int, ch chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	ch <- f
}

func main() {
	ch := make(chan int)

	for i := 5; i <= 15; i++ {
		go func(n int) {
			factorial(n, ch)
		}(i)

		fmt.Printf("%d! = %d\n", i, <-ch)
	}
}
