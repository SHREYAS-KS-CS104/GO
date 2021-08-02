package main

import (
	"fmt"
)

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	var a, b int
	fmt.Println("Enter two numbers: ")
	fmt.Scanf("%d%d", &a, &b)
	a, b = swap(a, b)
	fmt.Println("The Number entered is ", a, b)
}
