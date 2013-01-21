package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["key1"] = 23
	m["key2"] = 34
	fmt.Println(m)

	v1 := m["key1"]
	v2 := m["key2"]

	fmt.Println("v1: ", v1, ", v2: ", v2)

	delete(m, "key1")

	_, isPresent := m["key1"]
	fmt.Println("isPresent: ", isPresent)

	m1 := map[string]int{"key1": 2, "key2": 4}
	fmt.Println(m1)
}
