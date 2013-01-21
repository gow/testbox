package main

import "fmt"

func main() {
	var sum float32 = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
