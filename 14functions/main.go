package main

import "fmt"

func main() {
	fmt.Println("Functions in golang!")
	greeter()

	result := add(3, 5)
	fmt.Println("Result is:", result)

	proResult := proAdder(2, 4, 6, 8, 10)
	fmt.Println("Pro-Result is:", proResult)

	proResultAdv, msg := proMultiAdv(2, 4, 6, 8, 10)
	fmt.Println("Pro-Result Adv is:", proResultAdv)
	fmt.Println("Message is:", msg)
}

func greeter() {
	fmt.Println("Namaste from golang🙏")
}

func add(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func proMultiAdv(values ...int) (int, string) {
	total := 1
	for _, value := range values {
		total *= value
	}
	return total, "I'm advanced because I'm returing two different type values"
}
