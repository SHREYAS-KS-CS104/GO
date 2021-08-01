package main

import "fmt"

func gcd(m, n int) int {
	if m == n {
		return m
	} else if m > n {
		return gcd(m-n, n)
	} else {
		return gcd(m, n-m)
	}
}

func main() {
	var m, n int
	fmt.Print("Enter two numbers: ")
	fmt.Scanf("%d%d", &m, &n)
	fmt.Println("The GCD of the two number is: ", gcd(m, n))
}
