package main

import "fmt"

// Consider the money type declared at exercise #1.
// Create a new method for the money type called printStr
// that returns the money value as a string with 2 decimal points.

type money float64

// 2. Create a method called print that prints out the money value with exact 2 decimal points.
func (m money) printStr() string {
	return fmt.Sprintf("%.2f\n", m)
}

func main() {
	var m money = 1.127
	fmt.Println(m.printStr())
}
