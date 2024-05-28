package main

import "fmt"

func main() {
	fmt.Println("Methods")
	book1 := Book{"The White Tiger", true, 3, "Aravind Adiga", 43.2}
	fmt.Println(book1)
	book1.GetBookName()
	fmt.Println(book1.SetAvailableStatus())
	fmt.Println(book1)
}

type Book struct {
	Name        string
	IsAvailable bool
	Count       int
	Author      string
	Price       float32
}

func (b Book) GetBookName() {
	fmt.Println(b.Name)
}

func (b Book) SetAvailableStatus() bool {
	b.IsAvailable = false
	return b.IsAvailable
}
