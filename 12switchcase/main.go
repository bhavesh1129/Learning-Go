package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang!")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of your dice is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can open as your dice value 1")
	case 2:
		fmt.Println("You can move to 2 spots")
	case 3:
		fmt.Println("You can move to 3 spots")
		fallthrough
	case 4:
		fmt.Println("You can move to 4 spots")
		fallthrough
	case 5:
		fmt.Println("You can move to 5 spots")
	case 6:
		fmt.Println("You can move to 6 spots and roll dice again!")
	default:
		fmt.Println("What was that?")
	}
}
