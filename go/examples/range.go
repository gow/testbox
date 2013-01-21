package main

import "fmt"

func main() {
	num_arr := []int{2, 3, 4}
	var sum int = 0
	for _, num := range num_arr {
		sum += num
	}
	fmt.Println("Sum: ", sum)

	for i, num := range num_arr {
		if num == 3 {
			fmt.Println(num, "is present at index: ", i)
		}
	}

	kv_map := map[string]string{"a": "apple", "b": "ball"}
	for key, val := range kv_map {
		fmt.Printf("%s -> %s\n", key, val)
	}

	// Range over unicode code points.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
