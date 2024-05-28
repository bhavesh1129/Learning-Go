package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thrusday", "Friday"}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		if days[i] == "Monday" {
			goto greetMsg
		} else {
			fmt.Println(days[i])
		}
	}

greetMsg:
	fmt.Println("Hey! Good morning its monday")
}
