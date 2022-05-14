package main

import "fmt"

func main() {
	defer fmt.Println("Statement1")
	defer fmt.Println("Statement2")
	defer fmt.Println("Statement3")

	fmt.Println("Defers in gloang!")
	defersInGo()
	fmt.Println()
}

func defersInGo() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i, " ")
	}
}
