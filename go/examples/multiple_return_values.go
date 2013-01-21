package main

import "fmt"

func add_mul(a int, b float64) (int, float64) {
	return (a + int(b)), (float64(a) * b)
}

func main() {
	add_val, mul_val := add_mul(3, 4.5)
	fmt.Println("3 + 4.5 = ", add_val, "[int]")
	fmt.Println("3 * 4.5 = ", mul_val, "[float]")
}
