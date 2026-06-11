package main

import "fmt"

// Change the function from the previous exercise and use a `naked return`.
func sum(n ...int) (total int) {
	total = 0
	for _, num := range n {
		total += num
	}
	return
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	fmt.Println(sum())
}
