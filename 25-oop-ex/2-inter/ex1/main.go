package main

import "fmt"

// Consider the following type and interface declaration.

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

// 1. Create a Go program where car type implements the vehicle interface.

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

func main() {
	// 2. Create a variable of type vehicle that holds a car struct value.
	var v vehicle = car{
		licenceNo: "l1",
		brand:     "car1",
	}

	// 3. Call the methods (Licence and Name) on the interface value declared at step 2
	fmt.Println(v.License())
	fmt.Println(v.Name())
}
