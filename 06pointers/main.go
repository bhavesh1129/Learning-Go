package main

import "fmt"

func main() {
	fmt.Println("Intro of Pointers")

	var ptr *int //default value of pointer is nil
	fmt.Println(ptr)

	temp := 1129
	ptr = &temp
	fmt.Println("value of ptr:", ptr)
	fmt.Println("address of ptr:", *ptr)
}
