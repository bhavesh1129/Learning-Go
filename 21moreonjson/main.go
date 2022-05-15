package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //create alias
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON in golang!")
	encodeJSON()
}

func encodeJSON() {
	lcoCourse := []course{
		{
			"ReactJS Bootcamp", 299, "lco.in", "pass123", []string{
				"web-dev", "js", "react"},
		},
		{
			"MERN Bootcamp", 499, "lco.in",
			"pass456", []string{
				"mern", "web", "fullstack"},
		},
		{
			"Golang Bootcamp", 399, "lco.in", "pass789", nil,
		},
	}

	//package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", finalJson)
}
