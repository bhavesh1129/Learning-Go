package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.outgray.com/"

func main() {
	fmt.Println("Handling web requests")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T", response)

	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataBytes))
}
