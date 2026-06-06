package main

// Using Iota declare the following months of the year: Jun, Jul and Aug
func main() {
	//Jun, Jul and Aug are constant and their value is 6, 7 and 8
	const (
		Jun = iota + 6
		Jul
		Aug
	)
}
