package main

import "fmt"

func main() {
	fmt.Println("Methods in golang!")

	user := User{
		"Bhavesh Garg",
		"developerbhaveshgarg@gmail.com",
		true,
		20,
	}
	user.GetStatus()
	fmt.Println("Age of user is:", user.Age)
	fmt.Println("Updated age of user is:", user.GetAgePlusOne())
	user.NewMail()
	fmt.Println(user.Email)
}

//Since copy of value is passed to the methods hence original values can't be changed
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (user User) GetStatus() {
	fmt.Println("Is user active:", user.Status)
}

func (user User) GetAgePlusOne() int {
	return user.Age + 1
}

func (user User) NewMail() {
	user.Email = "bhaveshgarg151@gmail.com"
	fmt.Println("Updated Email of this user is:", user.Email)
}
