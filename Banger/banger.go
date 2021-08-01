package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	s = strings.Repeat("!", len(os.Args[1])) + strings.ToUpper(os.Args[1]) + strings.Repeat("!", len(os.Args[1]))
	//This program takes an input from the terminal and captitalizes it and add exclamation marks
	fmt.Println(s)
}
