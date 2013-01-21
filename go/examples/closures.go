package main

import "fmt"

func intSeq(start int) func() int {
	i := start
	return func() int {
		// every instance of this closeure encloses a separate copy of i which lives as long as the
		// closure
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq(-1)

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	var newInts = intSeq(55)
	fmt.Println(newInts())
	fmt.Println(newInts())

	i := 5
	var qqq = func() int {
		i += 2
		return i
	}

	fmt.Println(qqq())
	fmt.Println(qqq())
}
