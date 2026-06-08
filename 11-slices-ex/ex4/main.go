package main

import (
	"fmt"
	"os"
	"strconv"
)

// Create a Go program that reads some numbers from the command line
// and then calculates the sum and the product
// of all the numbers given at the command line.

func main() {
	sum, prod := 0., 1.

	for i := 1; i < len(os.Args); i++ {
		n, err := strconv.ParseFloat(os.Args[i], 64)
		if err != nil {
			continue
		}

		if n < 2 || n > 10 {
			fmt.Println("Please enter between 2 and 10 numbers!")
			return
		}

		sum += n
		prod *= n
	}

	fmt.Printf("sum = %.2f, prod = %.2f\n", sum, prod)
}
