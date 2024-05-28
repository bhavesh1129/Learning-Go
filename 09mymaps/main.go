package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)
	fmt.Println(languages)

	languages["js"] = "javascript"
	languages["js"] = "javascript"
	languages["jv"] = "java"
	languages["rb"] = "ruby"
	fmt.Println(languages)

	fmt.Println(len(languages))

	fmt.Println(languages["js"])

	delete(languages, "rb")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("%v -> %v\n", key, value)
	}
}
