package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{7, 2, 4}
	strings := []string{"c", "a", "b"}

	sort.Ints(ints)
	fmt.Println(ints, sort.IntsAreSorted(ints))
	sort.Strings(strings)
	fmt.Println(strings, sort.StringsAreSorted(strings))
}
