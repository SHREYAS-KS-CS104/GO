package main

import "fmt"

func main() {
	//Int go we can't declare a float64 value to int variable
	length := 100
	breath := 25.12
	fmt.Println("The area of the rectangle is: ", length*int(breath))
	//When we type cast a float64 to a int, it's decimal portion is truncated
	fmt.Println("The actual area of the rectangle is: ", float64(length)*breath)
	fmt.Println("The value of lenght before conversion is: ", length, "It's value after conversion is: ", float64(length))
	//Whenever we are type casting we have to close attention to which variable we want to type cast
}
