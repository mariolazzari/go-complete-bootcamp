package main

import "fmt"

// Declare two defined types called mile and kilometer.
// Have the underlying type be an float64.
type mile float64
type kilometer float64

// Declare a constant called m2km equals 1.609  ( 1mile=1.609km)
const m2km = 1.609

func main() {
	// declare a variable called mileBerlinToParis of type mile with value 655.3
	var mileBerlinToParis mile = 6555.3

	// declare a variable called kmBerlinToParis of type kilometer
	var kmBerlinToParis kilometer

	// calculate the distance between Berlin and Paris in km by multiplying mileBerlinToParis and m2km.
	// Assign the result to kmBerlinToParis
	kmBerlinToParis = kilometer(mileBerlinToParis * m2km)

	// print out the distance in km between Berlin and Paris
	fmt.Println(kmBerlinToParis)
}
