package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Create JSON Data")

	EncodeJSONData()

	DecodeJSONData()
}

type Book struct {
	Name        string `json:"bookname"`
	IsAvailable bool
	Count       int    `json:"numOfBooks"`
	Author      string `json:"-"`
	Price       float32
}

func EncodeJSONData() {
	books := []Book{
		{
			"To Kill a Mockingbird",
			true,
			12,
			"Harper Lee",
			10.99,
		},
		{
			"1984",
			false,
			0,
			"George Orwell",
			8.99,
		},
		{
			"The Great Gatsby",
			true,
			5,
			"F. Scott Fitzgerald",
			12.99,
		},
	}

	//package above data as json
	finalJson, err := json.MarshalIndent(books, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", finalJson)
	fmt.Printf("%T", finalJson)
}

func DecodeJSONData() {
	jsonDataFromWeb := []byte(`
		{
			"bookname": "To Kill a Mockingbird",
			"IsAvailable": true,
			"numOfBooks": 12,
			"Author": "Harper Lee",
			"Price": 10.99
		}
	`)

	var book Book
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON data is valid")
		json.Unmarshal(jsonDataFromWeb, &book)
		fmt.Printf("%#v\n", book)
	} else {
		fmt.Println("JSON data is not valid")
	}
}
