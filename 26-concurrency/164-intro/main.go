package main

import (
	"fmt"
	"time"
)

func doSomething() {
	fmt.Println("Do something...")
}

func main() {
	doSomething()
	go doSomething()

	time.Sleep(2 * time.Second)
}
