package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://outgray.com/our-work?client_id=11231&order_id=14278&order_id=12278"

func main() {
	fmt.Println("Handling URL")
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.RawPath)
	fmt.Println(result.Port())

	query := result.Query()
	fmt.Println(query)
	for k, v := range query {
		fmt.Println(k, " -> ", v)
	}
	fmt.Println(query["order_id"])
}
