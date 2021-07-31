package main

import (
	"fmt"
	"path"
)

func main() {
	dir, file := path.Split("Okay/super.pdf")
	fmt.Println("The file in the directory is ", dir, "and the file is ", file)
}
