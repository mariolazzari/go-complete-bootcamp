package main

import (
	"fmt"
	"time"
)

func main() {
	// unbuffered
	ch1 := make(chan int)
	// buffered
	ch2 := make(chan int, 3)
	_ = ch2

	go func(c chan int) {
		fmt.Println("Start goroutine")
		fmt.Println("After sending to channel")
		c <- 10 // blocks main until receving from channel
		fmt.Println("End goroutine")
	}(ch1)

	fmt.Println("Main sleeps")
	time.Sleep(time.Second)

	n := <-ch1
	fmt.Println("Main from channel:", n)
}
