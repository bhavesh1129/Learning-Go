package main

import "fmt"

func main() {
	fmt.Println("Defer")

	defer fmt.Println(0)
	fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
	defer fmt.Println(5)

	myDefer()
}

func myDefer() {
	for i := 6; i <= 10; i++ {
		defer fmt.Println(i)
	}
}
