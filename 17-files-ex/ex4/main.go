package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create a Go Program that reads the entire contents of a file called info.txt into a string.
	// You can use ioutil.ReadAll() or any other function you want.
	// The file exists in the current working directory.

	file, err := os.OpenFile("README.md", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	fmt.Println(string(fileBytes))
}
