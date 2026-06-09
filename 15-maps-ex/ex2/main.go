package main

import "fmt"

func main() {
	// There are some errors
	// var m1 map[int]bool
	var m1 = make(map[int]bool)
	m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	// There are some errors
	// fmt.Println(m2 == m3)
	m2Str := fmt.Sprintf("%v", m2)
	m3Str := fmt.Sprintf("%v", m3)
	fmt.Println(m2Str == m3Str)
}
