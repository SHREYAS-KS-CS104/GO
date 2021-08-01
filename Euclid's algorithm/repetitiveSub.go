package main

import "fmt"

func gcd(m, n int) int {
	for m != n {
		if m > n {
			m = m - n
		} else {
			n = n - m
		}
	}
	return m
}
func main() {
	var m, n int
	fmt.Print("Enter two numbers: ")
	fmt.Scanf("%d%d", &m, &n)
	fmt.Println("The GCD is: ", gcd(m, n))
}
