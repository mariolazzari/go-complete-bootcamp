package main

import "fmt"

func main() {
	// Consider the following map declaration:
	// m := map[int]bool{"1": true, 2: false, 3: false}

	// 1. The above map declaration has an error. Solve the error!
	m := map[int]bool{1: true, 2: false, 3: false}

	// 2. Delete a key:value pair from the map.
	delete(m, 1)

	// 3. Iterate over the map and print out each key and value.
	for k, v := range m {
		fmt.Printf("%d -> %t\n", k, v)
	}
}
