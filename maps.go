package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println(m)

	var v1 int
	var prs bool
	v1, prs = m["k1"]
	fmt.Println(prs, ":", v1)

	delete(m, "k2")
	fmt.Println(m)

	_, prs2 := m["k2"]
	fmt.Println(prs2)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}
