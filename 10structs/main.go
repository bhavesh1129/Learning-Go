package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang!")

	//no inheritance, no super, no parent-child relationship

	user := User{
		"Bhavesh Garg",
		"developerbhaveshgarg@gmail.com",
		true,
		20,
	}
	fmt.Println(user)
	fmt.Printf("Bhavesh Details are: %+v\n", user)
}

//Here User acts as class so, it should be capitalize and the fileds like Name, Email, Status, Age are also captialize because these values can be accessed from anywhere in the program and thats why we capitalize the fields (good practice)

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
