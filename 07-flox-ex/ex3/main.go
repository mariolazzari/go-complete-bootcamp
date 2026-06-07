package main

import "fmt"

// Change the code from the previous exercise and use the break statement to print out only the first 3 numbers divisible by 7 between 1 and 50.
func main() {
	j := 1
	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Println(i)

		j++
		if j == 4 {
			break
		}
	}
}
