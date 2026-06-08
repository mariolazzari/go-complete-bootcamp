package main

import "fmt"

func main() {
	// Consider the following slice declaration:
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}

	// Using a slice expression and a for loop iterate over the slice ignoring the first and the last two elements.
	var sum int

	// Print those elements and their sum.
	for _, num := range nums[1 : len(nums)-2] {
		fmt.Println(num)
		sum += num
	}
	fmt.Println("Sum:", sum)
}
