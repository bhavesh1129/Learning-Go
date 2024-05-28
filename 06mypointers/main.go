package main

import "fmt"

func main()  {
	fmt.Println("Pointers")

	// var ptr *int
	// myNum := 23
	// fmt.Println(ptr)
	// ptr = &myNum
	// fmt.Println(&ptr)

	// *ptr = *ptr * 2
	// fmt.Println(myNum)

	var ptr *int
	myNum := 34
	fmt.Printf("%T", ptr)
	ptr = &myNum
	fmt.Println(*ptr)
	fmt.Println(ptr)
	fmt.Println(&myNum)

}