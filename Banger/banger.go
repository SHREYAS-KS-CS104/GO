package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	excla := strings.Repeat("!", len(os.Args[1]))
	s = excla + os.Args[1] + excla
	//This program takes an input from the terminal and captitalizes it and add exclamation marks before and after
	fmt.Println(s)
}
