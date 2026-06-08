package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Printf("%#v\n", nums)

	nums = append(nums, 4)
	fmt.Printf("%#v\n", nums)

	nums = append(nums, 5, 6, 7)
	fmt.Printf("%#v\n", nums)

	nums2 := []int{100, 200}
	nums = append(nums, nums2...)
	fmt.Printf("%#v\n", nums)

	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	fmt.Printf("src = %#v, dst = %#v, nn = %d\n\n", src, dst, nn)

}
