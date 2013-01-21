package main

import "fmt"

func main() {
	s := make([]string, 4)
	fmt.Println("Empty slice: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	fmt.Println("Filled slice: ", s)

	fmt.Println("Slice length: ", len(s))

	s = append(s, "e")
	s = append(s, "f", "g")
	fmt.Println("Slice after appending: ", s)

	c := make([]string, 2)
	copy(c, s)
	fmt.Println("Copied slice: ", c)

	fmt.Println("Slice between 2 and 5: ", s[2:5])
	fmt.Println("Slice upto 5: ", s[:5])
	fmt.Println("Slice 3 onwards: ", s[3:])

	var l = []string{"aaa", "bbb", "ccc"}
	fmt.Println("Initialized slice: ", l)

	twoDslice := make([][]int, 3)
	for i := 0; i < len(twoDslice); i++ {
		innerLen := i + 1
		twoDslice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDslice[i][j] = i + j
		}
		fmt.Println("twoDslice[", i, "]", twoDslice[i])
	}
	fmt.Println("twoDslice: ", twoDslice)
}
