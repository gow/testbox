package main

import "fmt"

func main() {
	var a [5]int
	b := []int{1, 2, 3, 4, 5, 6}
	c := []string{"abc", "123", "!@#$%^&*()\"\"''", "'single quotes'"}

	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	var twoD [3][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("twoD=", twoD)
	fmt.Println("twoD len=", len(twoD))
}
