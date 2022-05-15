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
	// encodeJSON()
	decodeJson()
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

func decodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"Price": 299,
			"website": "lco.in",
			"tags": ["web-dev","js","react"]
		}
	`)

	// var lcoCourse course
	// checkValidJson := json.Valid(jsonDataFromWeb) //return bool value

	// if checkValidJson {
	// 	fmt.Println("JSON is valid!")
	// 	json.Unmarshal(jsonDataFromWeb, &lcoCourse)
	// fmt.Printf("Type of lcoCourse is: %T\n", lcoCourse)
	// 	fmt.Printf("%#v\n", lcoCourse)
	// } else {
	// 	fmt.Println("Json isn't valid")
	// }

	//Some cases where we want to add data to a key-value pair

	var myOnlineData map[string]interface{} //use interface because we don't the kind of data coming from web so we use interface, and rest it automatically handles

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	fmt.Printf("Type of myOnlineData is: %T\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key is %v and value is %v", key, value)
		fmt.Println()
	}
}
