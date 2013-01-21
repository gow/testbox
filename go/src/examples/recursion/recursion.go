package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return (n * factorial(n-1))
}

func main() {
	fmt.Printf("Factorial of %d is %d\n", 10, factorial(10))
	fmt.Printf("Factorial of %d is %d\n", 3, factorial(3))
	fmt.Printf("Factorial of %d is %d\n", 5, factorial(5))
	fmt.Printf("Factorial of %d is %d\n", 0, factorial(0))
	fmt.Printf("Factorial of %d is %d\n", 1, factorial(1))
}
