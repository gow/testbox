package main

import "fmt"

func add(a int, b int) int {
	return (a + b)
}

func main() {
	fmt.Println("3 + 6 =", add(3, 6))
}
