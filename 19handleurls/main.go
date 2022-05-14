package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=golang&paymentid=bhavesh1129"

func main() {
	fmt.Println("Handle URL's in golang!")

	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Result is", result.Scheme)
	fmt.Println("Result is", result.Host)
	fmt.Println("Result is", result.Path)
	fmt.Println("Result is", result.RawQuery)
	fmt.Println("Result is", result.Port())

	queryparams := result.Query()
	fmt.Println("Query Params:", queryparams)
	fmt.Printf("Type of Query Params is %T\n", queryparams)
	fmt.Println(queryparams["paymentid"])

	for key, value := range queryparams {
		fmt.Println("Key:", key+" ,Value:", value)
	}

	//Reverse Logic
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=bhavesh",
	}
	fmt.Println(partsOfUrl)
	resString:=partsOfUrl.String()
	fmt.Println(resString)
}
