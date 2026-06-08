package main

import "fmt"

func main() {
	// 1. Declare a slice called nums with three float64 numbers.
	nums := []float64{1.1, 2.2, 3.3}
	// 2. Append the value 10.1 to the slice
	nums = append(nums, 1.1)
	// 3. In one statement append to the slice the values: 4.1,  5.5 and 6.6
	nums = append(nums, 4.1, 5.5, 6.6)
	// 4. Print out the slice
	fmt.Printf("%#v\n", nums)
	// 5. Declare a slice called n with two float64 values
	n := []float64{7.7, 8.8}
	// 6. Append n to nums
	nums = append(nums, n...)
	// 7. Print out the nums slice
	fmt.Printf("%#v\n", nums)
}
