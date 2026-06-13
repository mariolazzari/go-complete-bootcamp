package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		fmt.Printf("%s is dow\n", url)
		return
	}

	defer resp.Body.Close()
	fmt.Printf("%s -> %d\n", url, resp.StatusCode)
	if resp.StatusCode == 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		file := strings.Split(url, "//")[1]
		fmt.Printf("Writing response to %s\n", file)
		err = os.WriteFile(file, body, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	urls := []string{"https://mariolazzari.it", "https://mariafilippini.it", "https://golang.org"}
	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
	}

	fmt.Printf("Goroutines running: %d\n", runtime.NumGoroutine())

	wg.Wait()
}
