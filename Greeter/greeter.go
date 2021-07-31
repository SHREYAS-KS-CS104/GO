package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	if name[0] > 'a' && name[0] < 'z' {
		fmt.Println("You used a lowercase letter")
	}
	fmt.Print(" /\\_/\\\n( o.o )\n > ^ <\n")
	fmt.Printf("Hello %s, This is Ms. kitty\n", name)

}
