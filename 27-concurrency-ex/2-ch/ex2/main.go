package main

import "fmt"

// Create a function literal (a.k.a. anonymous function)
// that sends the string value if receives as argument
// to main func using a channel.
func main() {
	ch := make(chan string)

	go func() {
		ch <- "Ciao"
	}()

	fmt.Println(<-ch)
}
