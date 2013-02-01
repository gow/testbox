package main

import (
	"fmt"
	"sort"
)

func DemoSort() {
	strs := []string{"qww", "afa", "q", "aaa"}
	sort.Strings(strs)
	fmt.Println("Sorted strings: ", strs)

	ints := []int{2, 5, 3, 7, 1, 8}
	sort.Ints(ints)
	fmt.Println("Sorted ints: ", ints)

	is_sorted := sort.IntsAreSorted(ints)
	fmt.Println("Is Sorted? ", is_sorted)
}

type Date struct {
	day   int
	month int
	year  int
}

type DateArr []Date

func (d DateArr) Len() int {
	return len(d)
}
func (d DateArr) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d DateArr) Less(i, j int) bool {
	if d[i].year == d[j].year {
		if d[i].month == d[j].month {
			return d[i].day < d[j].day
		}
		return d[i].month < d[j].month
	}
	return d[i].year < d[j].year
}

func DemoCustomSort() {
	dates := DateArr{{20, 10, 1982}, {7, 5, 1983}, {20, 10, 1982}, {20, 10, 1986}, {20, 10, 1984}}
	fmt.Println(dates)
	sort.Sort(dates)
	fmt.Println(dates)
}

func main() {
	DemoSort()
	DemoCustomSort()
}
