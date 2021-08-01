package main

import (
	"fmt"
)

func main() {
	var num1, num2, r int
	fmt.Print("Enter two number: ")
	fmt.Scanf("%d%d", &num1, &num2)
	for num2 != 0 {
		r = num1 % num2
		num1 = num2
		num2 = r
	}
	fmt.Println("The GCD of the two numbers is: ", num1)
}
