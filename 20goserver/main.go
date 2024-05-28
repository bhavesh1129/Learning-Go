package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Go server")
	PerformGetRequest()

	PerformPostRequest()

	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl string = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	fmt.Println(response.ContentLength)

	// result, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(result))

	var resFromString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := resFromString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(resFromString.String())
}

func PerformPostRequest() {
	const myurl string = "http://localhost:8000/post"

	//sending fake json data
	reqBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":12,
			"platform":"www.google.com"
		}
	`)
	res, err := http.Post(myurl, "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl string = "http://localhost:8000/postform"

	//send fake form data
	data := url.Values{}
	data.Add("firstname", "bhavesh")
	data.Add("lastname", "garg")
	data.Add("email", "bhavesh@gmail.com")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}
