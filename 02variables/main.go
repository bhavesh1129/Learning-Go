package main

import "fmt"

const LoginToken string = "abcd" //public access due to first capital letter

func main() {
	var username string = "bhavesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("Variable is of type: %T \n", isLogged)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	//default values and aliases
	var anotherVar string
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	//implicit type
	var website = "folware.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	//no var style
	numOfUser := 100000
	fmt.Println(numOfUser)
	fmt.Printf("Variable is of type: %T \n", numOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}