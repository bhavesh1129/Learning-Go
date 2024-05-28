package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

var signals = []string{"test"}

var mut sync.Mutex

func main() {
	fmt.Println("Go routines")
	// go greeter("Hello")
	// greeter("world")

	websiteList := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.apple.com",
		"https://www.microsoft.com",
	}

	for _, website := range websiteList {
		go GetStatusCode(website)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func greeter(s string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func GetStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
	}
	fmt.Printf("%d %s\n", res.StatusCode, endpoint)
}
