package main

import (
	"log"
	"os"
)

func main() {
	// Rename the file created at Exercise #1 to data.txt
	// Check if file exists before renaming it.
	// If it doesn't exist print a message and stop the program.

	oldName := "info.txt"
	newName := "data.txt"

	_, err := os.Stat(oldName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("File %s does not exist\n", oldName)
		}
	}

	err = os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}
