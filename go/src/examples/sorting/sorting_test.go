package main

import "sort"
import "fmt"

func ExampleSorting() {
	dates := DateArr{{20, 10, 1982}, {7, 5, 1983}, {20, 10, 1982}, {20, 10, 1986}, {20, 10, 1984}}
	sort.Sort(dates)
	fmt.Println(dates)
	// Output:
	// [{20 10 1982} {20 10 1982} {7 5 1983} {20 10 1984} {20 10 1986}]
}
