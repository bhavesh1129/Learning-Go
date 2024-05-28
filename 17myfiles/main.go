package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files")

	content := "I am learning about files in golang"
	file, err := os.Create("./dummyText.txt")
	checkNilErr(err)

	defer file.Close()

	length, err := WriteFile(file, content)
	checkNilErr(err)
	fmt.Println("length is ", length)

	readFile, err := ReadFile("./dummyText.txt")
	checkNilErr(err)
	fmt.Println(string(readFile))
}

func ReadFile(fileName string) (string, error) {
	readFile, err := os.ReadFile(fileName)
	return string(readFile), err
}

func WriteFile(file *os.File, content string) (int, error) {
	length, err := io.WriteString(file, content)
	return length, err
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
