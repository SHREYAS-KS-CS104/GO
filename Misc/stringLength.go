package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var name string
	name = "Shreyas"
	fmt.Println(utf8.RuneCountInString(name))
	name = "Ộzil"
	fmt.Println(utf8.RuneCountInString(name))

}
