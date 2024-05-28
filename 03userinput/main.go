package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	reader:=bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	fmt.Println("Type a number")
	fmt.Printf(input)
}