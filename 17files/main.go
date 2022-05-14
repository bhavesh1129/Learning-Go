package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in golang!")

	content := "This needs to go in file -hello.com"
	fmt.Println(content)

	// for creating, majority of time we use os package
	file, err := os.Create("./myTempFile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	// for writing, majority of time we use io package
	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is:", length)

	readFile("./myTempFile.txt")
	defer file.Close()
}

func readFile(filename string) {

	// for reading, majority of time we use ioutil package
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)
	content := string(databyte)
	fmt.Println("Content in file is:", content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
