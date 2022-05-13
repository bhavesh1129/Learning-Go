package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in golang!")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	//Method-1
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	//Method-2
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	//Method-3
	// for _, value := range days {
	// 	fmt.Println(value)
	// }

	//Similar to while loop
	tempValue := 1
	for tempValue <= 10 {

		if tempValue == 7 {
			break
		}

		//goto
		if tempValue == 2 {
			goto flag
		}

		if tempValue == 5 {
			tempValue++
			continue
		}

		fmt.Println("Value is:", tempValue)
		tempValue++
	}

flag:
	fmt.Println("Jumping to flag")
}
