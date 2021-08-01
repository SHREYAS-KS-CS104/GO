package main

import "fmt"

func small(m, n int) int {
	if m < n {
		return m
	} else {
		return n
	}
}

func gcd(m, n int) int {
	min := small(m, n)
	for m%min != 0 || n%min != 0 {
		min -= 1
	}
	return min
}

func main() {
	var m, n int
	fmt.Print("Enter two numbers: ")
	fmt.Scanf("%d%d", &m, &n)
	fmt.Println("The GCD of the two numbers is: ", gcd(m, n))
}
