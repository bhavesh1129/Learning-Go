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

	file, err := os.Create("./myTempFile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is:", length)

	readFile("./myTempFile.txt")
	defer file.Close()
}

func readFile(filename string) {
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
