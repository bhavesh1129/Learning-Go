package main

import "fmt"

func main() {
	fmt.Println("If else in golang!")

	loginCount := 15
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 && loginCount <= 20 {
		result = "Gold user"
	} else {
		result = "Premium user"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Required result is even")
	} else {
		fmt.Println("Required result is odd")
	}

	if num := 13; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number isn't less than 10")
	}
}
