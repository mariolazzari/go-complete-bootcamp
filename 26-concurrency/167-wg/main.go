package main

import (
	"fmt"
	"runtime"
	"sync"
)

func f1(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("f1() start")
	for i := range 3 {
		fmt.Println("f1() i =", i)
	}
	fmt.Println("f1() end")
}

func f2(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("f2() start")
	for i := range 8 {
		fmt.Println("f2() i =", i)
	}
	fmt.Println("f2() end")
}

func main() {
	fmt.Println("Main execution started")

	var wg sync.WaitGroup

	wg.Add(1)
	go f1(&wg)
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	wg.Add(1)
	go f2(&wg)
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("Main execution ended")
}
