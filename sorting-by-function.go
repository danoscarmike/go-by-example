package main

import (
	"fmt"
	"sort"
)

type byLength []string
type reverse []int

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (r reverse) Len() int {
	return len(r)
}

func (r reverse) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r reverse) Less(i, j int) bool {
	return r[j] < r[i]
}

func main() {
	fruits := []string{"pomegranate", "apple", "banana"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	digits := []int{1, 27, 42, 16}
	sort.Sort(reverse(digits))
	fmt.Println(digits)
}
