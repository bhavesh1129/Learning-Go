package main

import "fmt"

func main() {
	fmt.Println("Maps in golang!")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)
	fmt.Printf("Type of language is %T\n", languages)

	fmt.Println("JS shorts for:", languages["JS"])

	//how to delete items in the map in golang
	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)

	//loops in golang
	for key, value := range languages {
		fmt.Println(key + " " + value)

		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	//comma ok syntax in loops
	for _, value := range languages {
		fmt.Println(value)

		fmt.Printf("value is %v\n", value)
	}
}
