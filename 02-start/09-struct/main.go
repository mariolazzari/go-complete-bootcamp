package main

import "fmt"

func main() {
	fmt.Println("Ciao!")
	dist := 60.8
	fmt.Printf("%.2f Km => %.2f miles\n", dist, dist*0.621)
}
