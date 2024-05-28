package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time")
	presentTime := time.Now()
	fmt.Println(presentTime)

	 createdTime := time.Date(2024, time.December, 31, 24, 56, 59, 59, time.UTC)
	 fmt.Println(createdTime.Format("01-02-2006 Monday"))
}
