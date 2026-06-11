package main

import "fmt"

// Create a Go program with a function called f1() that takes a parameter of type uint and returns 2 values:
// a) the factorial of n
// b) the sum of all integer numbers greater than zero (>0) and less than or equal to n (<=n)

func f1(n uint) (uint, uint) {
	var fact uint = 1
	var sum uint

	// factorial of n
	var i uint = n
	for i > 0 {
		fact *= i
		i--
	}

	// sum of [0,n]
	i = 0
	for i <= n {
		sum += i
		i++
	}

	return fact, sum
}

func main() {
	f, s := f1(1)
	fmt.Printf("fact = %d, sum = %d\n", f, s)
	f, s = f1(2)
	fmt.Printf("fact = %d, sum = %d\n", f, s)
	f, s = f1(3)
	fmt.Printf("fact = %d, sum = %d\n", f, s)
	f, s = f1(4)
	fmt.Printf("fact = %d, sum = %d\n", f, s)
}
