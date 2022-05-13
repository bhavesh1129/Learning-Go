package main

import "fmt"

func main() {
	fmt.Println("Intro to Arrays")

	// var myList [4]int
	// myList[0] = 1
	// myList[1] = 2
	// myList[2] = 3
	// myList[3] = 4
	// fmt.Println("myList is:", myList)

	var myList [4]string
	myList[0] = "1"
	myList[1] = "2"
	// myList[2] = 3
	myList[3] = "4"
	fmt.Println("myList is:", myList)

	var newList = [4]int{1, 2, 3, 4}
	fmt.Println("newList is:", newList)
}
