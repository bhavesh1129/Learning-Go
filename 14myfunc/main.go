package main

import "fmt"

func main() {
	fmt.Println("Functions")
	message := greetMsg()
	fmt.Println(message)

	func() {
		fmt.Println("IIFI!")
	}()

	sum, msg := proAdder(2, 3, 4, 1, 5, 6, 3, 2, 2, 2)
	fmt.Println(sum, msg)

	dummy()
}

func dummy() {
	fmt.Println("hii")
}

func proAdder(v ...int) (int, string) {
	totalSum := 0
	for _, v := range v {
		totalSum += v
	}
	return totalSum, "ProAdder"
}

func greetMsg() string {
	return ("Good morning!")
}
