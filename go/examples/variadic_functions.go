package main

import "fmt"

func sum(nums ...int) (total int64) {
	fmt.Println("Number of inputs:", len(nums))
	for _, num := range nums {
		total += int64(num)
	}
	return total
}

func main() {
	fmt.Println(sum(1, 3))
	fmt.Println(sum(1, 3, 4))

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(sum(nums[:3]...)) // converts a slice into multiple args
}
