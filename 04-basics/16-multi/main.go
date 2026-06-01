package main

import "fmt"

func main() {
	car, price := "Aygo", 20_000
	fmt.Printf("%s price: %d\n", car, price)

	var (
		myCar     = "Yaris"
		myPrice   float64
		available bool
	)

	myPrice = 30000
	available = true

	if available {
		fmt.Printf("%s price: %.2f\n", myCar, myPrice)
	}

}
