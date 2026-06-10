package main

import (
	"log"
	"os"
)

func main() {
	// Create a new file in the current working directory called info.txt and then close the file.
	// If the file cannot be created (no permissions, wrong path etc) then print out the error and stop the program (do error handling).
	file, err := os.Create("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
