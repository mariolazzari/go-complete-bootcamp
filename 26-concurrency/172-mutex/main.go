package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	const gr = 100
	var wg sync.WaitGroup

	wg.Add(gr * 2)

	var n int = 0

	// starting 200 goroutines with mutex
	var mx sync.Mutex

	for range gr {
		go func() {
			defer wg.Done()
			time.Sleep(time.Second / 10)

			mx.Lock()
			n++
			mx.Unlock()
		}()

		go func() {
			defer wg.Done()
			time.Sleep(time.Second / 10)

			mx.Lock()
			defer mx.Unlock()
			n--
		}()

	}
	wg.Wait()
	fmt.Println(n)
}
