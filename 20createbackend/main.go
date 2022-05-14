package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	fmt.Println("Create backend using golang!")

	//GET REQUETS
	// PerformGetRequest()

	//POST REQUETS -JSON Data
	// PerformPostJsonRequest()

	//POST REQUETS -Form Data
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	checkNilErr(err)

	defer response.Body.Close()

	fmt.Println("Status Code of request:", response.StatusCode)
	fmt.Println("Content Length is:", response.ContentLength)

	var responseString strings.Builder
	databyte, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	// fmt.Println(string(databyte))

	byteCount, err := responseString.Write(databyte)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//Fake JSON Payload(sending some data)
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":100,
			"platform":"lco",
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	checkNilErr(err)

	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println(string(databyte))

	file, err := os.Create("./postedData.txt")
	checkNilErr(err)

	io.WriteString(file, string(databyte))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//Form Data
	data := url.Values{}
	data.Add("firstname", "bhavesh")
	data.Add("lastname", "garg")
	data.Add("email", "bhaveshgarg151@gmail.com")

	response, err := http.PostForm(myurl, data)
	checkNilErr(err)

	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println(string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
