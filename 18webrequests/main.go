package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web-requets in golang!")

	response, err := http.Get(url)
	checkNilErr(err)

	// fmt.Println("Response is:", response)
	fmt.Printf("Response is of type %T\n:", response)

	//caller's responsibility to close the connection
	defer response.Body.Close()

	databyte, _ := ioutil.ReadAll(response.Body)

	file, err := os.Create("./fetchedData.txt")
	checkNilErr(err)

	content, err := io.WriteString(file, string(databyte))
	checkNilErr(err)
	fmt.Println(content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
