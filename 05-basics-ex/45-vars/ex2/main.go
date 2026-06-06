package main

func main() {
	// Declare a, b, c, d using a single var keyword (multiple variable declaration) for better readability.
	var (
		a int
		b float64
		c bool
		d string
	)

	// Declare x, y and z on a single line -> multiple short declarations
	x, y, z := 20, 15.5, "Gopher!"

	// Remove the statement that prints out the variables. See the error!
	//fmt.Println(a, b, c, d, y, x, z)

	// Change the program to run without error using the blank identifier
	_ = a
	_ = b
	_ = c
	_ = d
	_ = x
	_ = y
	_ = z
}
