package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
	// "math/rand"
)

func main() {
	fmt.Println("Maths in Golang")

	// var numOne int = 2
	// var numTwo float32 = 4.5
	// fmt.Println(numOne + int(numTwo))

	// num:= rand.Intn(10)
	// fmt.Println(num)

	// num:=rand.Uint32()
	// fmt.Println(num)

	num, _ := rand.Int(rand.Reader, big.NewInt(time.Now().UnixNano()))
	fmt.Println(num)
}
