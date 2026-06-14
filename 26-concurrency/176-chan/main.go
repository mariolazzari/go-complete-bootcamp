package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n
}

func main() {
	var ch chan int
	fmt.Println(ch)

	ch = make(chan int)
	fmt.Println(ch)

	// send operation
	go func() {
		ch <- 42
	}()

	// receive operation
	n := <-ch

	fmt.Println("n =", n)
	// fmt.Println("n =", <-ch)

	// send only
	c1 := make(chan<- string)
	// receive only
	c2 := make(<-chan string)

	fmt.Printf("%T, %T, %T\n", ch, c1, c2)

	// goroutines and channels
	go f1(50, ch)
	n = <-ch
	fmt.Println("n from f1 ch:", n)

	// close channel
	close(ch)
}
