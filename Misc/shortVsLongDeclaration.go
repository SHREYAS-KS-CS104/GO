package main

import "fmt"

var global int = 50

func main() {
	fmt.Println("\n\n\nThe value of the global variable is: ", global)
	global := 90
	fmt.Println("The value of global now is: ", global)
	{
		//This a new block
		global := 4000
		fmt.Println("The value inside the new block is: ", global)
	}
	//Since, the scope of a variable is limitted to it's block the i value remains unchanged outside
	fmt.Println("The value of i outside is still: ", global)
}
