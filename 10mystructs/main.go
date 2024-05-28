package main

import "fmt"

func main()  {
	fmt.Println("Structs")

	dummyUser := User{"dummyUser", true, 22}
	fmt.Println(dummyUser)

	fmt.Printf("%+v\n", dummyUser)

	fmt.Printf("%v", dummyUser.Name)


}

type User struct{
Name string
IsAlive bool
Age int
}