package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1() start")
	for i := range 3 {
		fmt.Println("f1() i =", i)
	}
	fmt.Println("f1() end")
}

func f2() {
	fmt.Println("f1() start")
	for i := range 8 {
		fmt.Println("f1() i =", i)
	}
	fmt.Println("f1() end")
}

func main() {
	fmt.Println("Main execution started")

	fmt.Println("Logical CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("Max Procs:", runtime.GOMAXPROCS(0))

	go f1()
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	go f2()
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	time.Sleep(time.Millisecond * 200)
	fmt.Println("Main execution ended")
}
