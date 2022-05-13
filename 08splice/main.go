package main

import (
	"fmt"
	// "sort"
)

func main() {
	fmt.Println("Intro to splice")

	// var list = []string{"a", "b", "c"}
	// fmt.Println("list is:", list)
	// fmt.Printf("Type of list is %T\n", list)

	// list = append(list, "d", "e", "f")
	// fmt.Println("list is:", list)

	// list = append(list[1:3])
	// fmt.Println("list is:", list)

	// list = append(list[:3])
	// fmt.Println("list is:", list)

	//Unique feature of Golang
	// highScores := make([]int, 4)
	// highScores[0] = 234
	// highScores[1] = 789
	// highScores[2] = 456
	// highScores[3] = 901
	// highScores[4] = 441

	// highScores = append(highScores, 555, 435, 777)
	// fmt.Println("Highscores is:", highScores)

	// fmt.Println("isHighScoresSorted: ", sort.IntsAreSorted(highScores))
	// sort.Ints(highScores)
	// fmt.Println("Sorted highscores:", highScores)
	// fmt.Println("isHighScoresSorted: ", sort.IntsAreSorted(highScores))

	//Remove a value from slices based on index
	var courses = []string{"java", "javascript", "c++", "python", "go"}
	fmt.Println("Courses:", courses)

	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("updated courses:", courses)
}
