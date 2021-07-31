package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string
	dir, file = path.Split("Woah/cool.pdf")
	fmt.Println("Directory: ", dir)
	fmt.Println("File: ", file)
}
