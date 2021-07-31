package main

import (
	"fmt"
	"path"
)

func main() {
	var file string
	//Here i'm using a blank identifier to discard the values
	_, file = path.Split("Example/woah.pdf")
	fmt.Println("The file is: ", file)
}
