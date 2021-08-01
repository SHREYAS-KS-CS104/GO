package main

import "fmt"

func gcd(m, n int) int {
	if n == 0 {
		return m
	} else {
		return gcd(n, m%n)
	}
}
func main() {
	var m, n int
	fmt.Print("Enter two number: ")
	fmt.Scanf("%d%d", &m, &n)
	ans := gcd(m, n)
	fmt.Println("The gcd of the two numbers is: ", ans)
}
