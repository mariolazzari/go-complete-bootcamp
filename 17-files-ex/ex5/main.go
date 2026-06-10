package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a Go Program that reads the entire contents of a file called info.txt  using a scanner (bufio package) line by line.
	// The file exists in the current working directory.

	file, err := os.OpenFile("README.md", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
