package main

import (
	"fmt"
	"sync"
)

func sayHello(n string) {
	fmt.Printf("Hello, %s!\n", n)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		sayHello("Mr. Wick")
	}()

	wg.Wait()
}
