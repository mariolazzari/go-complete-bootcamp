package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "Ciao"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "Hello"
	}()

	time.Sleep(time.Second)

	// wait from multiple channels
	for range 2 {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		default:
			fmt.Println("No activity")
		}
	}

	fmt.Printf("Exec time: %v\n", time.Since(start))
}
