package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial value: ", i)

	zeroVal(i)
	fmt.Println("After zeroVal: ", i)

	zeroPtr(&i)
	fmt.Println("After zeroPtr: ", i)

	fmt.Println("Pointer: ", &i)
}
