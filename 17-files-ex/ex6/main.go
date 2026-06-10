package main

import (
	"io/ioutil"
	"log"
)

func main() {
	// 1. Using single function creates a file called info.txt in the current directory.
	// If the file already exists it will truncate it to zero size.

	// file, err := os.Create("info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// 2. Write the string "The Go gopher is an iconic mascot!" to the file.
	// _, err = file.WriteString("The Go gopher is an iconic mascot!")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	bs := []byte("The Go gopher is an iconic mascot!")
	err := ioutil.WriteFile("info.txt", bs, 0644)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}
