package main

import "fmt"

func zeroval(n int) {
	n = 0
}

func zeroptr(n *int) {
	*n = 0
}

func main() {
	n := 1
	fmt.Println("initial:", n)

	zeroval(n)
	fmt.Println("zeroval:", n)

	zeroptr(&n)
	fmt.Println("zeroptr:", n)

	fmt.Println("pointer:", &n)
}
