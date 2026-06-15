package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	go func(c chan int) {
		for i := range 5 {
			fmt.Printf("Start goroutine %d\n", i+1)
			fmt.Println("After sending to channel")
			c <- i + 1
			fmt.Printf("End goroutine %d\n", i+1)
		}
		close(c) // all data sent
	}(ch)

	//fmt.Println("Main sleeps")
	//time.Sleep(time.Second)

	// receive fron data sending channel
	for v := range ch { // v:=<-c
		fmt.Println("Main received:", v)
	}

	// zero value after closing
	fmt.Println("Main closed ch:", <-ch)

	// c <-10 panic!
}
