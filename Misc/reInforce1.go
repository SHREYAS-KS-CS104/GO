package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("The length of the name is: ", len(os.Args[1]))
	color := "green"
	color = "Dark " + "bule"
	n := 0.
	pi := 3.14
	num := 2
	n = float64(num) * pi
	fmt.Println("The value of n is: ", n)
	fmt.Println("The color is: ", color)
	width, length := 6, 5
	fmt.Println("The perimeter of the rectangle is: ", 2*(length+width))
}
