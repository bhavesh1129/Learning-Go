package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	// var slice = []string{}
	// fmt.Println(len(slice))

	// slice = append(slice, "hello", "hi")
	// fmt.Println(len(slice))

	// slice = append(slice, "hola", "hey")
	// fmt.Println(len(slice))

	// fmt.Println(slice)

	slices2 := make([]string, 4)
	slices2[0]="H"
	slices2[1]="E"
	slices2[2]="L"
	slices2[3]="L"
	fmt.Println(slices2)

	slices2 = append(slices2, "hola", "hey", "hello", "hi")
	fmt.Println(slices2)

	isSorted := sort.StringsAreSorted(slices2)
	fmt.Println(isSorted)
	// sort.Strings(slices2)
	// fmt.Println(slices2)

	sort.Sort(sort.Reverse(sort.StringSlice(slices2)))
	fmt.Println(slices2)


	var courses = []string{"Maths", "Science", "English", "Hindi", "GK"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
