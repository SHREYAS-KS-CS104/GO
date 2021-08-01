package main

import "fmt"

func main() {
	const (
		monday = iota
		_      //The Blank identifier consumes the iota
		tuesday
		wednesday
		thrusday
		friday
		saturday
		sunday
	)
	fmt.Println(monday, tuesday, wednesday, thrusday, friday, saturday, sunday)
}
