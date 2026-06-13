package main

import "fmt"

// 1. Create a new struct type called book with 2 fields:
// price and title of type float64 and string.
type book struct {
	price float64
	title string
}

// 2. Create a method for the newly defined type called
// vat that returns the vat value if vat is 9%.
func (b book) vat() float64 {
	return b.price * 0.09
}

// 3. Create a method called discount that takes a book value as receiver and
// decreases the price of the book by 10%.
func (b book) discount() float64 {
	return b.price * 0.9
}

func main() {
	book := book{title: "The Trial", price: 9.9}
	fmt.Printf("Vat:%v\n", book.vat())
	fmt.Printf("%#v\n", book.discount())
}
