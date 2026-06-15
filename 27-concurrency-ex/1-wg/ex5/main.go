package main

import (
	"fmt"
	"sync"
)

func deposit(b *int, n int, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()

	mx.Lock()
	defer mx.Unlock()
	*b += n
	//wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()

	mx.Lock()
	defer mx.Unlock()
	*b -= n
	// wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex

	wg.Add(200)

	balance := 100

	for i := range 100 {
		go deposit(&balance, i, &wg, &mx)
		go withdraw(&balance, i, &wg, &mx)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}
