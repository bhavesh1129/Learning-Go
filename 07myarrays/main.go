package main

import "fmt"

func main()  {
	fmt.Println("Arrays")

	// var array [5]string
	// array[0]="1"

	var array2 = [4]string{"hello", "hii", "hola", "namaste"}
	fmt.Println(len(array2))
	fmt.Println(array2[1])
}